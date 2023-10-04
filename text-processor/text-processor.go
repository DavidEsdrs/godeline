package text_processor

import (
	"bytes"
	"fmt"
	"regexp"

	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/logger"
	"github.com/DavidEsdrs/goditor/position"
	"github.com/DavidEsdrs/goditor/tags"
	"github.com/DavidEsdrs/goditor/token"
	"github.com/DavidEsdrs/goditor/tracker"
)

type Processor struct {
	EditionTree     *editnode.EditionTree
	logger          *logger.Logger
	maxBufferLength int
}

func NewProcessor(et *editnode.EditionTree, maxBufferLength int, logger *logger.Logger) Processor {
	return Processor{
		EditionTree:     et,
		logger:          logger,
		maxBufferLength: maxBufferLength,
	}
}

// Tokenize the given text using the processor prefix tree (EditionTree)
func (p *Processor) Tokenize(text string, sanitize bool) TextResult {
	currentIdx := 0
	textLength := len(text)
	result := NewTextResult()

	tracker := tracker.NewTracker()

	for currentIdx < textLength {
		alreadySeen := tracker.AlreadySeen(currentIdx)

		if !alreadySeen {
			tag, found := p.FoundTag(text, currentIdx)

			if found {
				token, _ := p.GetTextByTag(text, currentIdx, tag) // TODO: Handle error properly
				result.AddToken(&token)
				tracker.RegisterToken(token)
			}
		}

		currentIdx++
	}

	if sanitize {
		for _, t := range result.tokens {
			t.Word = Normalize(t.Word, t.Tag.Opening, t.Tag.Closing)
		}
	}

	return result
}

// returns the text within the given tags
func (p *Processor) GetTextByTag(text string, idx int, tag tags.Tag) (token.Token, error) {
	var result token.Token

	startingIdx := idx
	// since this function is being called when we've found a opening tag, we assume
	// that the next few characters is the given opening tag. So let skip it by
	// offsetting it
	offset := len(tag.Opening)
	currentIdx := idx + offset
	bufferLen := len(tag.Closing)
	startingPosition := position.GetPosition(text, idx)
	currentCol := startingPosition.Col
	currentLn := startingPosition.Ln
	textLength := len(text)

	found := false

	for !found && currentIdx-startingIdx < p.maxBufferLength && currentIdx+bufferLen < textLength {
		buffer := text[currentIdx : currentIdx+bufferLen]

		if buffer == tag.Closing {
			innerText := text[startingIdx : currentIdx+bufferLen]
			pos := position.Position{
				Ln:    currentLn,
				Col:   currentCol + offset - len(innerText),
				Index: idx,
			}
			result = token.NewToken(innerText, pos, tag, nil) // TODO: pass editNode instead of nil
			return result, nil
		}

		currentIdx, currentCol, currentLn = updatePosition(text, currentIdx, currentCol, currentLn)
	}

	return result, fmt.Errorf("couldn't find the correspondent closing tag")
}

func updatePosition(text string, currentIdx, currentCol, currentLn int) (int, int, int) {
	currentChar := rune(text[currentIdx])

	if currentChar == '\n' {
		currentLn++
		currentCol = 0
	} else {
		currentCol++
	}

	currentIdx++

	return currentIdx, currentCol, currentLn
}

// returns wether the next few characters constitute a tag
func (p *Processor) FoundTag(text string, idx int) (tags.Tag, bool) {
	current := p.EditionTree.Root()
	currentIdx := idx
	currentChar := rune(text[currentIdx])

	node, exists := current.Children[currentChar]

	// if exists is true, there is a possibility that we have found a tag
	// we will just make sure when continue to read the next few characters
	for exists {
		current = node
		currentIdx++

		currentChar = rune(text[currentIdx])
		node, exists = current.Children[currentChar]
	}

	return current.Tag, current.IsEnd
}

func (p *Processor) Sanitize(tokens ...token.Token) {
	for i, t := range tokens {
		sanitizedText := Normalize(t.Word, t.EditNode.Tag.Opening, t.EditNode.Tag.Closing)
		tokens[i].Word = sanitizedText
	}
}

// create a new string from the source with the given segments removed
func Normalize(source string, removable ...string) string {
	var buf bytes.Buffer
	for _, str := range removable {
		escapedStr := regexp.QuoteMeta(str)
		if buf.Len() > 0 {
			buf.WriteByte('|')
		}
		buf.WriteString(escapedStr)
	}
	regex := regexp.MustCompile(buf.String())
	output := regex.ReplaceAllString(source, "")
	return output
}
