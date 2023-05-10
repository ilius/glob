package match

import (
	"fmt"
	"strings"
)

type Suffix struct {
	Suffix string
}

func NewSuffix(s string) Suffix {
	return Suffix{s}
}

func (suf Suffix) Len() int {
	return lenNo
}

func (suf Suffix) Match(s string) bool {
	return strings.HasSuffix(s, suf.Suffix)
}

func (suf Suffix) Index(s string) (int, []int) {
	idx := strings.Index(s, suf.Suffix)
	if idx == -1 {
		return -1, nil
	}

	return 0, []int{idx + len(suf.Suffix)}
}

func (suf Suffix) String() string {
	return fmt.Sprintf("<suffix:%s>", suf.Suffix)
}
