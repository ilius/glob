package match

import (
	"fmt"
	"unicode/utf8"
)

type Min struct {
	Limit int
}

func NewMin(l int) Min {
	return Min{l}
}

func (mn Min) Match(s string) bool {
	var l int
	for range s {
		l += 1
		if l >= mn.Limit {
			return true
		}
	}

	return false
}

func (mn Min) Index(s string) (int, []int) {
	var count int

	c := len(s) - mn.Limit + 1
	if c <= 0 {
		return -1, nil
	}

	segments := acquireSegments(c)
	for i, r := range s {
		count++
		if count >= mn.Limit {
			segments = append(segments, i+utf8.RuneLen(r))
		}
	}

	if len(segments) == 0 {
		return -1, nil
	}

	return 0, segments
}

func (mn Min) Len() int {
	return lenNo
}

func (mn Min) String() string {
	return fmt.Sprintf("<min:%d>", mn.Limit)
}
