package token

import (
	editnode "github.com/DavidEsdrs/godeline/edit-node"
	"github.com/DavidEsdrs/godeline/position"
	"github.com/DavidEsdrs/godeline/tags"
)

type Token struct {
	// Holds the inner text
	InnerText string
	// Position holds the starting position of the token
	Position position.Position
	Length   int
	// by now, has no use (always nil)
	EditNode *editnode.EditNode
	Tag      tags.Tag
}

func NewToken(word string, Position position.Position, Tag tags.Tag, EditNode *editnode.EditNode) Token {
	return Token{
		InnerText: word,
		Position:  Position,
		Length:    len(word),
		Tag:       Tag,
		EditNode:  EditNode,
	}
}

func (t Token) IsEmpty() bool {
	return t.InnerText != "" && t.Length > 0
}
