package match

import (
	"fmt"
	"strings"

	sutil "github.com/gobwas/glob/util/strings"
)

type SuffixAny struct {
	Suffix     string
	Separators []rune
}

func NewSuffixAny(s string, sep []rune) SuffixAny {
	return SuffixAny{s, sep}
}

func (sa SuffixAny) Index(s string) (int, []int) {
	idx := strings.Index(s, sa.Suffix)
	if idx == -1 {
		return -1, nil
	}

	i := sutil.LastIndexAnyRunes(s[:idx], sa.Separators) + 1

	return i, []int{idx + len(sa.Suffix) - i}
}

func (sa SuffixAny) Len() int {
	return lenNo
}

func (sa SuffixAny) Match(s string) bool {
	if !strings.HasSuffix(s, sa.Suffix) {
		return false
	}
	return sutil.IndexAnyRunes(s[:len(s)-len(sa.Suffix)], sa.Separators) == -1
}

func (sa SuffixAny) String() string {
	return fmt.Sprintf("<suffix_any:![%s]%s>", string(sa.Separators), sa.Suffix)
}
