package token

import (
	editnode "github.com/DavidEsdrs/goditor/editNode"
	"github.com/DavidEsdrs/goditor/position"
	"github.com/DavidEsdrs/goditor/tags"
)

type Token struct {
	Word     string
	Position position.Position
	Length   int
	EditNode *editnode.EditNode
	Tag      tags.Tag
}

func NewToken(word string, Position position.Position, Tag tags.Tag, EditNode *editnode.EditNode) Token {
	return Token{
		Word:     word,
		Position: Position,
		Length:   len(word),
		Tag:      Tag,
		EditNode: EditNode,
	}
}

func (t Token) IsEmpty() bool {
	return t.Word != "" && t.Length > 0
}
