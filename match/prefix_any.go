package match

import (
	"fmt"
	"strings"
	"unicode/utf8"

	sutil "github.com/gobwas/glob/util/strings"
)

type PrefixAny struct {
	Prefix     string
	Separators []rune
}

func NewPrefixAny(s string, sep []rune) PrefixAny {
	return PrefixAny{s, sep}
}

func (pa PrefixAny) Index(s string) (int, []int) {
	idx := strings.Index(s, pa.Prefix)
	if idx == -1 {
		return -1, nil
	}

	n := len(pa.Prefix)
	sub := s[idx+n:]
	i := sutil.IndexAnyRunes(sub, pa.Separators)
	if i > -1 {
		sub = sub[:i]
	}

	seg := acquireSegments(len(sub) + 1)
	seg = append(seg, n)
	for i, r := range sub {
		seg = append(seg, n+i+utf8.RuneLen(r))
	}

	return idx, seg
}

func (pa PrefixAny) Len() int {
	return lenNo
}

func (pa PrefixAny) Match(s string) bool {
	if !strings.HasPrefix(s, pa.Prefix) {
		return false
	}
	return sutil.IndexAnyRunes(s[len(pa.Prefix):], pa.Separators) == -1
}

func (pa PrefixAny) String() string {
	return fmt.Sprintf("<prefix_any:%s![%s]>", pa.Prefix, string(pa.Separators))
}
