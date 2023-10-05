package text_processor

import (
	"fmt"

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
	stopOnError     bool
}

func NewProcessor(et *editnode.EditionTree, maxBufferLength int, logger *logger.Logger) Processor {
	return Processor{
		EditionTree:     et,
		logger:          logger,
		maxBufferLength: maxBufferLength,
		stopOnError:     false,
	}
}

func (p *Processor) StopOnError() {
	p.stopOnError = true
}

var accGettingPosition int64 = 0

// Tokenize the given text using the processor prefix tree (EditionTree)
func (p *Processor) Tokenize(text string, sanitize bool) (TextResult, error) {
	textLength := len(text)
	result := NewTextResult()

	currentPosition := position.Position{Ln: 0, Col: 0, Index: 0}

	tracker := tracker.NewTracker()

	for idx := 0; idx < textLength; idx++ {
		if !tracker.AlreadySeen(idx) {
			tag, found := p.FoundTag(text, idx)

			if found {
				token, err := p.GetTextByTag(text, idx, tag, currentPosition)

				if err != nil && p.stopOnError {
					return result, err
				}

				if err == nil {
					result.AddToken(&token)
					tracker.RegisterToken(token)
				}
			}
		}

		currentPosition.Index, currentPosition.Col, currentPosition.Ln = updatePosition(text, idx, currentPosition.Col, currentPosition.Ln)
	}

	if sanitize {
		sanitizeTokens(result.tokens)
	}

	fmt.Printf("accGettingPosition: %v\n", accGettingPosition)

	return result, nil
}

// returns the text within the given tags
func (p *Processor) GetTextByTag(text string, idx int, tag tags.Tag, startingPosition position.Position) (token.Token, error) {
	var result token.Token

	startingIdx := idx
	// since this function is being called when we've found a opening tag, we assume
	// that the next few characters is the given opening tag. So let skip it by
	// offsetting it
	offset := len(tag.Opening)
	currentIdx := idx + offset
	bufferLen := len(tag.Closing)
	currentCol := startingPosition.Col
	currentLn := startingPosition.Ln
	textLength := len(text)

	for currentIdx-startingIdx < p.maxBufferLength && currentIdx+bufferLen < textLength {
		buffer := text[currentIdx : currentIdx+bufferLen]

		if buffer == tag.Closing {
			innerText := text[startingIdx : currentIdx+bufferLen]
			pos := position.Position{
				Ln:    currentLn,
				Col:   currentCol,
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

func sanitizeTokens(tokens []*token.Token) {
	for i := range tokens {
		tokens[i].Word = Normalize(tokens[i].Word, tokens[i].Tag.Opening, tokens[i].Tag.Closing)
	}
}
