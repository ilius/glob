package match

import (
	"fmt"
	"unicode/utf8"
)

type Max struct {
	Limit int
}

func NewMax(l int) Max {
	return Max{l}
}

func (mx Max) Match(s string) bool {
	var l int
	for range s {
		l += 1
		if l > mx.Limit {
			return false
		}
	}

	return true
}

func (mx Max) Index(s string) (int, []int) {
	segments := acquireSegments(mx.Limit + 1)
	segments = append(segments, 0)
	var count int
	for i, r := range s {
		count++
		if count > mx.Limit {
			break
		}
		segments = append(segments, i+utf8.RuneLen(r))
	}

	return 0, segments
}

func (mx Max) Len() int {
	return lenNo
}

func (mx Max) String() string {
	return fmt.Sprintf("<max:%d>", mx.Limit)
}
