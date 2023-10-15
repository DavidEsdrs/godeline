package tracker

import (
	"github.com/DavidEsdrs/godeline/token"
)

// Tracker is used to keep track of the closing tags position.
// It is util for avoiding re-reading of closing tags, which could lead to errors
// when reading tags with equal opening and closing tags.
type Tracker struct {
	tagsAlreadySeen map[int]struct{}
}

func NewTracker() Tracker {
	return Tracker{
		tagsAlreadySeen: make(map[int]struct{}),
	}
}

// RegisterToken should be used register a new closing tag position
func (t *Tracker) RegisterToken(token token.Token) {
	closingTagIndex := token.Position.Index + token.Length - len(token.Tag.Closing)
	t.tagsAlreadySeen[closingTagIndex] = struct{}{}
}

func (t *Tracker) RegisterIndex(idx int) {
	t.tagsAlreadySeen[idx] = struct{}{}
}

// this function allow us to check if we already saw a position
func (t *Tracker) AlreadySeen(idx int) bool {
	_, exists := t.tagsAlreadySeen[idx]
	return exists
}

func (t *Tracker) Clear() {
	t.tagsAlreadySeen = map[int]struct{}{}
}
