package godeline

import (
	"errors"

	editnode "github.com/DavidEsdrs/godeline/edit-node"
	"github.com/DavidEsdrs/godeline/position"
	"github.com/DavidEsdrs/godeline/tags"
	text_processor "github.com/DavidEsdrs/godeline/text-processor"
	"github.com/DavidEsdrs/godeline/token"
	"github.com/DavidEsdrs/godeline/tracker"
)

var (
	ErrClosingTagNotFound = errors.New("godeline: source code has not closing tag")
)

type Processor struct {
	EditionTree     *editnode.EditionTree
	maxBufferLength int
	stopOnError     bool
	sanitize        bool
}

func NewProcessor(et *editnode.EditionTree, maxBufferLength int) Processor {
	return Processor{
		EditionTree:     et,
		maxBufferLength: maxBufferLength,
		stopOnError:     false,
	}
}

func (p *Processor) StopOnError() {
	p.stopOnError = true
}

func (p *Processor) Sanitize() {
	p.sanitize = true
}

// Tokenize the given text using the processor prefix tree (EditionTree)
func (p *Processor) Tokenize(text string) (text_processor.TextResult, error) {
	textLength := len(text)
	result := text_processor.NewTextResult()

	currentPosition := position.Position{Ln: 0, Col: 0, Index: 0}

	tracker := tracker.NewTracker()

	for idx := 0; idx < textLength; idx++ {
		if !tracker.AlreadySeen(idx) {
			tag, found := p.foundTag(text, textLength, idx)

			if found {
				token, err := p.getTextByTag(text, idx, tag, currentPosition, tracker)

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

	return result, nil
}

// returns the text within the given tags
func (p *Processor) getTextByTag(text string, idx int, tag tags.Tag, startingPosition position.Position, tracker tracker.Tracker) (token.Token, error) {
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
	openingLen := len(tag.Opening)

	for canSearch(currentIdx, startingIdx, p.maxBufferLength, textLength, bufferLen) {
		if !tracker.AlreadySeen(currentIdx) {
			buffer := text[currentIdx : currentIdx+bufferLen]

			if buffer == tag.Closing {
				var innerText string

				if p.sanitize {
					innerText = text[startingIdx+openingLen : currentIdx]
				} else {
					innerText = text[startingIdx : currentIdx+bufferLen]
				}

				pos := position.Position{
					Ln:    currentLn,
					Col:   currentCol,
					Index: idx,
				}
				result = token.NewToken(innerText, pos, tag, nil) // TODO: pass editNode instead of nil
				return result, nil
			}
		}

		currentIdx, currentCol, currentLn = updatePosition(text, currentIdx, currentCol, currentLn)
	}

	return result, ErrClosingTagNotFound
}

func canSearch(currIdx, startingIdx, maxLength, textLength, bufferLen int) bool {
	return currIdx < textLength && currIdx+bufferLen < textLength && (currIdx-startingIdx < maxLength || maxLength == 0)
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
func (p *Processor) foundTag(text string, textLength, idx int) (tags.Tag, bool) {
	current := p.EditionTree.Root()
	currentIdx := idx
	currentChar := rune(text[currentIdx])

	node, exists := current.Children[currentChar]

	// if exists is true, there is a possibility that we have found a tag
	// we will just make sure when continue to read the next few characters
	for exists {
		current = node
		currentIdx++

		if idx < textLength {
			currentChar = rune(text[currentIdx])
			node, exists = current.Children[currentChar]
		}
	}

	return current.Tag, current.IsEnd
}
