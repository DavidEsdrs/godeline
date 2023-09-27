package token

import (
	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/tags"
)

type Token struct {
	Word     string
	Col, Ln  int
	Length   int
	EditNode *editnode.EditNode
	Tag      tags.Tag
}

func NewToken(word string, Col, Ln int, EditNode *editnode.EditNode) Token {
	return Token{
		Word:     word,
		Col:      Col,
		Ln:       Ln,
		Length:   len(word),
		EditNode: EditNode,
	}
}
