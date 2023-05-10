package match

import (
	"fmt"
	"unicode/utf8"

	"github.com/gobwas/glob/util/runes"
)

// single represents ?
type Single struct {
	Separators []rune
}

func NewSingle(s []rune) Single {
	return Single{s}
}

func (sing Single) Match(s string) bool {
	r, w := utf8.DecodeRuneInString(s)
	if len(s) > w {
		return false
	}

	return runes.IndexRune(sing.Separators, r) == -1
}

func (sing Single) Len() int {
	return lenOne
}

func (sing Single) Index(s string) (int, []int) {
	for i, r := range s {
		if runes.IndexRune(sing.Separators, r) == -1 {
			return i, segmentsByRuneLength[utf8.RuneLen(r)]
		}
	}

	return -1, nil
}

func (sing Single) String() string {
	return fmt.Sprintf("<single:![%s]>", string(sing.Separators))
}
