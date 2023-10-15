package text_processor

import (
	"bytes"
	"regexp"

	"github.com/DavidEsdrs/godeline/token"
)

type Sanitizer struct {
}

// removes the opening and closing tags from a token
func (p Sanitizer) Sanitize(tokens ...token.Token) {
	for i, t := range tokens {
		sanitizedText := Normalize(t.Word, t.EditNode.Tag.Opening, t.EditNode.Tag.Closing)
		tokens[i].Word = sanitizedText
	}
}

// creates a new string from the source with the given segments removed
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

// Removes non-alpha-numeric characters from src string
func (p Sanitizer) RemoveNonAlpha(src string) string {
	regex := regexp.MustCompile("[^a-zA-Z]+")
	output := regex.ReplaceAllString(src, "")
	return output
}
