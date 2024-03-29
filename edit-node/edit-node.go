package editnode

import (
	"fmt"

	"github.com/DavidEsdrs/godeline/stack"
	"github.com/DavidEsdrs/godeline/tags"
)

type EditionTree struct {
	root       *EditNode
	NodesQuant int // TODO: Remove
}

func NewEditTree() EditionTree {
	node := newEditNode('0')
	return EditionTree{
		root:       &node,
		NodesQuant: 1,
	}
}

func (et *EditionTree) Root() *EditNode {
	return et.root
}

func (et *EditionTree) Traverse() []EditNode {
	result := []EditNode{}
	stack := stack.NewStack[EditNode]()
	stack.Push(*et.root)
	for stack.Length > 0 {
		current := stack.Pop()
		result = append(result, current)
		for _, c := range current.Children {
			stack.Push(*c)
		}
	}
	return result
}

func (et *EditionTree) Reset() {
	node := newEditNode('0')
	et.root = &node
}

func (et *EditionTree) AddDelimiterType(delimiter string, counterpart string) error {
	if !isValidDelimiter(delimiter) {
		return fmt.Errorf("invalid delimiter")
	}

	var current *EditNode = et.root

	for i := range delimiter {
		currSegment := rune(delimiter[i])
		node, found := current.Children[currSegment]

		if found {
			current = node
			continue
		}

		newNode := EditNode{
			Segment:  currSegment,
			Children: map[rune]*EditNode{},
			Tag:      tags.Tag{Opening: delimiter, Closing: counterpart},
			IsEnd:    i == len(delimiter)-1,
		}

		current.Children[currSegment] = &newNode
		current = &newNode

		et.NodesQuant++ // TODO: Remove

		found = false
	}

	return nil
}

func isValidDelimiter(delimiter string) bool {
	return len(delimiter) > 0
}

type EditNode struct {
	Segment  rune
	Tag      tags.Tag
	IsEnd    bool
	Children map[rune]*EditNode
}

func newEditNode(seg rune) EditNode {
	return EditNode{
		Segment:  seg,
		Children: map[rune]*EditNode{},
	}
}
