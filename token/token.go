package token

import editnode "github.com/DavidEsdrs/goditor/editNode"

type Token struct {
	Word     string
	Col, Ln  int
	Length   int
	EditNode *editnode.EditNode
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
