package match

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Prefix struct {
	Prefix string
}

func NewPrefix(p string) Prefix {
	return Prefix{p}
}

func (p Prefix) Index(s string) (int, []int) {
	idx := strings.Index(s, p.Prefix)
	if idx == -1 {
		return -1, nil
	}

	length := len(p.Prefix)
	var sub string
	if len(s) > idx+length {
		sub = s[idx+length:]
	} else {
		sub = ""
	}

	segments := acquireSegments(len(sub) + 1)
	segments = append(segments, length)
	for i, r := range sub {
		segments = append(segments, length+i+utf8.RuneLen(r))
	}

	return idx, segments
}

func (p Prefix) Len() int {
	return lenNo
}

func (p Prefix) Match(s string) bool {
	return strings.HasPrefix(s, p.Prefix)
}

func (p Prefix) String() string {
	return fmt.Sprintf("<prefix:%s>", p.Prefix)
}
