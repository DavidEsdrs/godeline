package text_processor

import (
	"regexp"

	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/logger"
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
	result := NewTextResult()

	accIndex := 0
	currLine := 0
	charsUntilLastLine := 0
	length := len(text)

	for accIndex < length {
		current := rune(text[accIndex])
		bufferLength := 1

		if current == '\n' {
			charsUntilLastLine = accIndex
			currLine++
		}

		if isSegmentStart, _ := p.EditionTree.IsSegmentStart(current); isSegmentStart {
			delimiter, foundDelimiter := p.EditionTree.GetFirstDelimiter(text[accIndex:])

			if foundDelimiter {
				counterpart := delimiter.DelimiterCounterpart
				finalText, matchedDelimiter := ReadUntilSegment(text[accIndex:], counterpart, p.maxBufferLength)
				bufferLength = len(delimiter.FullDelimiter)

				if matchedDelimiter {
					relativeIdx := accIndex - charsUntilLastLine - 1
					t := token.NewToken(finalText, relativeIdx, currLine, &delimiter)
					result.AddToken(t)
				}
			}
		}

		accIndex += bufferLength
	}

	if sanitize {
		result.tokens = p.Sanitize(result.tokens...)
		return result
	}

	return result
}

func (p *Processor) Sanitize(tokens ...token.Token) []token.Token {
	res := make([]token.Token, len(tokens))
	for i, t := range tokens {
		sanitizedText := Normalize(t.Word, t.EditNode.FullDelimiter, t.EditNode.DelimiterCounterpart)
		res[i] = t
		res[i].Word = sanitizedText
	}
	return res
}

// returns the whole string until find the given segment
func ReadUntilSegment(text, segment string, maxBufferLength int) (res string, foundSegment bool) {
	bufferLen := len(segment)
	var buffer string

	i := bufferLen

	for i+bufferLen < len(text) && i < maxBufferLength {
		buffer = text[i : i+bufferLen]
		if buffer == segment {
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
