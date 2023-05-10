package match

import (
	"fmt"
	"strings"
)

type Contains struct {
	Needle string
	Not    bool
}

func NewContains(needle string, not bool) Contains {
	return Contains{needle, not}
}

func (c Contains) Match(s string) bool {
	return strings.Contains(s, c.Needle) != c.Not
}

func (c Contains) Index(s string) (int, []int) {
	var offset int

	idx := strings.Index(s, c.Needle)

	if !c.Not {
		if idx == -1 {
			return -1, nil
		}

		offset = idx + len(c.Needle)
		if len(s) <= offset {
			return 0, []int{offset}
		}
		s = s[offset:]
	} else if idx != -1 {
		s = s[:idx]
	}

	segments := acquireSegments(len(s) + 1)
	for i := range s {
		segments = append(segments, offset+i)
	}

	return 0, append(segments, offset+len(s))
}

func (c Contains) Len() int {
	return lenNo
}

func (c Contains) String() string {
	var not string
	if c.Not {
		not = "!"
	}
	return fmt.Sprintf("<contains:%s[%s]>", not, c.Needle)
}
