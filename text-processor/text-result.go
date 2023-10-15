package text_processor

import "github.com/DavidEsdrs/goditor/token"

type TextResult struct {
	tokens        []*token.Token
	TokenQuantity int
}

func NewTextResult() TextResult {
	return TextResult{
		tokens: []*token.Token{},
	}
}

// Appends a new token in the array
func (tr *TextResult) AddToken(token *token.Token) {
	tr.tokens = append(tr.tokens, token)
	tr.TokenQuantity++
}

// Tokens returns the tokens added with AddToken
func (tr *TextResult) Tokens() []*token.Token {
	return tr.tokens
}

func (tr *TextResult) Stats() {}
