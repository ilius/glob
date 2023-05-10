package match

import (
	"fmt"
	"unicode/utf8"

	"github.com/gobwas/glob/util/runes"
)

type List struct {
	List []rune
	Not  bool
}

func NewList(list []rune, not bool) List {
	return List{list, not}
}

func (ls List) Match(s string) bool {
	r, w := utf8.DecodeRuneInString(s)
	if len(s) > w {
		return false
	}

	inList := runes.IndexRune(ls.List, r) != -1
	return inList == !ls.Not
}

func (ls List) Len() int {
	return lenOne
}

func (ls List) Index(s string) (int, []int) {
	for i, r := range s {
		if ls.Not == (runes.IndexRune(ls.List, r) == -1) {
			return i, segmentsByRuneLength[utf8.RuneLen(r)]
		}
	}

	return -1, nil
}

func (ls List) String() string {
	var not string
	if ls.Not {
		not = "!"
	}

	return fmt.Sprintf("<list:%s[%s]>", not, string(ls.List))
}
