package match

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// raw represents raw string to match
type Text struct {
	Str         string
	RunesLength int
	BytesLength int
	Segments    []int
}

func NewText(s string) Text {
	return Text{
		Str:         s,
		RunesLength: utf8.RuneCountInString(s),
		BytesLength: len(s),
		Segments:    []int{len(s)},
	}
}

func (tx Text) Match(s string) bool {
	return tx.Str == s
}

func (tx Text) Len() int {
	return tx.RunesLength
}

func (tx Text) Index(s string) (int, []int) {
	index := strings.Index(s, tx.Str)
	if index == -1 {
		return -1, nil
	}

	return index, tx.Segments
}

func (tx Text) String() string {
	return fmt.Sprintf("<text:`%v`>", tx.Str)
}
