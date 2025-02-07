package match

import (
	"fmt"
	"unicode/utf8"
)

type Range struct {
	Lo, Hi rune
	Not    bool
}

func NewRange(lo, hi rune, not bool) Range {
	return Range{lo, hi, not}
}

func (Range) Len() int {
	return lenOne
}

func (rng Range) Match(s string) bool {
	r, w := utf8.DecodeRuneInString(s)
	if len(s) > w {
		return false
	}

	inRange := r >= rng.Lo && r <= rng.Hi

	return inRange == !rng.Not
}

func (rng Range) Index(s string) (int, []int) {
	for i, r := range s {
		if rng.Not != (r >= rng.Lo && r <= rng.Hi) {
			return i, segmentsByRuneLength[utf8.RuneLen(r)]
		}
	}

	return -1, nil
}

func (rng Range) String() string {
	var not string
	if rng.Not {
		not = "!"
	}
	return fmt.Sprintf("<range:%s[%s,%s]>", not, string(rng.Lo), string(rng.Hi))
}
