package editnode

import (
	"fmt"

	"github.com/DavidEsdrs/goditor/stack"
)

type EditionTree struct {
	root       *EditNode
	NodesQuant int // TODO: Remove
}

func NewEditTree() EditionTree {
	node := NewEditNode('0')
	return EditionTree{
		root:       &node,
		NodesQuant: 1,
	}
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

func (et *EditionTree) NewEditionType(delimiter string, counterpart string) error {
	if !isValidDelimiter(delimiter) {
		return fmt.Errorf("invalid delimiter")
	}

	var current *EditNode = et.root

	for i := range delimiter {
		currSegment := rune(delimiter[i])
		found := false

		for _, node := range current.Children {
			if node.Segment == currSegment {
				found = true
				current = node
				break
			}
		}

		if found {
			continue
		}

		newNode := EditNode{
			Segment:              currSegment,
			Children:             []*EditNode{},
			FullDelimiter:        delimiter,
			DelimiterCounterpart: counterpart,
			IsEnd:                i == len(delimiter)-1,
		}
		current.Children = append(current.Children, &newNode)
		current = &newNode

		et.NodesQuant++ // TODO: Remove

		found = false
	}

	return nil
}

func (et *EditionTree) IsSegmentStart(r rune) (isSegmentStart bool, editNode EditNode) {
	for _, node := range et.root.Children {
		if node.Segment == r {
			return true, *node
		}
	}
	return
}

// returns if the given text has any delimiter in it - if so, it returns true
// and the delimiter found
func (et *EditionTree) HasDelimiter(text string) (hasDelimiter bool, delimiter string) {
	current := et.root

	for _, r := range text {
		for _, c := range current.Children {
			if c.Segment == r {
				current = c
				break
			}
		}
	}

	hasDelimiter = current.Segment != '0' && current.IsEnd
	delimiter = current.FullDelimiter

	return
}

// receives a text and returns the first delimiter found within it
func (et *EditionTree) GetFirstDelimiter(text string) (node EditNode, found bool) {
	current := et.root

	for _, r := range text {
		if r == ' ' || r == '\n' {
			return *current, false
		}
		for _, c := range current.Children {
			if c.Segment == r {
				current = c
				if c.IsEnd {
					return *current, true
				}
				break
			}
		}
	}

	return
}

func isValidDelimiter(delimiter string) bool {
	return len(delimiter) > 0
}

type EditNode struct {
	Segment              rune
	FullDelimiter        string
	DelimiterCounterpart string
	IsEnd                bool
	Children             []*EditNode
}

func NewEditNode(seg rune) EditNode {
	return EditNode{
		Segment:  seg,
		Children: []*EditNode{},
	}
}
