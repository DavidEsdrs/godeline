package text_processor

import (
	"regexp"

	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/logger"
	"github.com/DavidEsdrs/goditor/position"
	"github.com/DavidEsdrs/goditor/tags"
	"github.com/DavidEsdrs/goditor/token"
)

type Processor struct {
	EditionTree     *editnode.EditionTree
	Tokens          []string
	logger          *logger.Logger
	maxBufferLength int
}

func NewProcessor(et *editnode.EditionTree, maxBufferLength int, logger *logger.Logger) Processor {
	return Processor{
		EditionTree:     et,
		Tokens:          []string{},
		logger:          logger,
		maxBufferLength: maxBufferLength,
	}
}

func (p *Processor) Tokenize(text string, sanitize bool) TextResult {
	currentIdx := 0
	textLength := len(text)
	result := NewTextResult()

	for currentIdx < textLength {
		tag, found := p.FoundTag(text, currentIdx)

		if found {
			token := p.GetTextByTag(text, currentIdx, tag)
			result.AddToken(token)
		}

		currentIdx++
	}

	return result
}

// returns the text within the given tags
func (p *Processor) GetTextByTag(text string, idx int, tag tags.Tag) token.Token {
	var result token.Token

	startingIdx := idx
	// once the function is being called when we've found a opening tag, we assume
	// that the next few characters is the given opening tag. So let skip it.
	offset := len(tag.Opening)
	currentIdx := idx + offset
	bufferLen := len(tag.Closing)
	startingPosition := position.GetPosition(text, idx)
	currentCol := startingPosition.Col
	currentLn := startingPosition.Ln

	found := false

	for !found && currentIdx-startingIdx < p.maxBufferLength {
		buffer := text[currentIdx : currentIdx+bufferLen]

		if buffer == tag.Closing {
			innerText := text[startingIdx : currentIdx+bufferLen]
			pos := position.Position{Ln: currentLn, Col: currentCol + offset - len(innerText)}
			result = token.NewToken(innerText, pos, tag, nil) // TODO: pass editNode instead of nil
			return result
		}

		currentIdx, currentCol, currentLn = updatePosition(text, currentIdx, currentCol, currentLn)
	}

	return result
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

func (p *Processor) Sanitize(tokens ...token.Token) []token.Token {
	res := make([]token.Token, len(tokens))
	for i, t := range tokens {
		sanitizedText := Normalize(t.Word, t.EditNode.Tag.Opening, t.EditNode.Tag.Closing)
		res[i] = t
		res[i].Word = sanitizedText
	}
	return res
}

// returns the whole string until find the given segment
func ReadUntilTag(text string, tag tags.Tag, maxBufferLength int) (res string, foundSegment bool) {
	bufferLen := len(tag.Closing)
	var buffer string

	i := bufferLen

	for i+bufferLen < len(text) && i < maxBufferLength {
		buffer = text[i : i+bufferLen]
		if buffer == tag.Closing {
			return text[:i+bufferLen], true
		}
		i++
	}

	return "", false
}

// create a new string from the source with the given segments removed
func Normalize(source string, removable ...string) string {
	pattern := ""
	for _, str := range removable {
		escapedStr := regexp.QuoteMeta(str)
		if pattern == "" {
			pattern = escapedStr
		} else {
			pattern = pattern + "|" + escapedStr
		}
	}
	regex := regexp.MustCompile(pattern)
	output := regex.ReplaceAllString(source, "")
	return output
}
