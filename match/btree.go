package match

import (
	"fmt"
	"unicode/utf8"
)

type BTree struct {
	Value            Matcher
	Left             Matcher
	Right            Matcher
	ValueLengthRunes int
	LeftLengthRunes  int
	RightLengthRunes int
	LengthRunes      int
}

func NewBTree(Value, Left, Right Matcher) (tree BTree) {
	tree.Value = Value
	tree.Left = Left
	tree.Right = Right

	lenOk := true
	if tree.ValueLengthRunes = Value.Len(); tree.ValueLengthRunes == -1 {
		lenOk = false
	}

	if Left != nil {
		if tree.LeftLengthRunes = Left.Len(); tree.LeftLengthRunes == -1 {
			lenOk = false
		}
	}

	if Right != nil {
		if tree.RightLengthRunes = Right.Len(); tree.RightLengthRunes == -1 {
			lenOk = false
		}
	}

	if lenOk {
		tree.LengthRunes = tree.LeftLengthRunes + tree.ValueLengthRunes + tree.RightLengthRunes
	} else {
		tree.LengthRunes = -1
	}

	return tree
}

func (bt BTree) Len() int {
	return bt.LengthRunes
}

// todo?
func (bt BTree) Index(s string) (index int, segments []int) {
	//inputLen := len(s)
	//// try to cut unnecessary parts
	//// by knowledge of length of right and left part
	//offset, limit := bt.offsetLimit(inputLen)
	//for offset < limit {
	//	// search for matching part in substring
	//	vi, segments := bt.Value.Index(s[offset:limit])
	//	if index == -1 {
	//		return -1, nil
	//	}
	//	if bt.Left == nil {
	//		if index != offset {
	//			return -1, nil
	//		}
	//	} else {
	//		left := s[:offset+vi]
	//		i := bt.Left.IndexSuffix(left)
	//		if i == -1 {
	//			return -1, nil
	//		}
	//		index = i
	//	}
	//	if bt.Right != nil {
	//		for _, seg := range segments {
	//			right := s[:offset+vi+seg]
	//		}
	//	}

	//	l := s[:offset+index]
	//	var left bool
	//	if bt.Left != nil {
	//		left = bt.Left.Index(l)
	//	} else {
	//		left = l == ""
	//	}
	//}

	return -1, nil
}

func (bt BTree) Match(s string) bool {
	inputLen := len(s)
	// try to cut unnecessary parts
	// by knowledge of length of right and left part
	offset, limit := bt.offsetLimit(inputLen)

	for offset < limit {
		// search for matching part in substring
		index, segments := bt.Value.Index(s[offset:limit])
		if index == -1 {
			releaseSegments(segments)
			return false
		}

		l := s[:offset+index]
		var left bool
		if bt.Left != nil {
			left = bt.Left.Match(l)
		} else {
			left = l == ""
		}

		if left {
			for i := len(segments) - 1; i >= 0; i-- {
				length := segments[i]

				var right bool
				var r string
				// if there is no string for the right branch
				if inputLen <= offset+index+length {
					r = ""
				} else {
					r = s[offset+index+length:]
				}

				if bt.Right != nil {
					right = bt.Right.Match(r)
				} else {
					right = r == ""
				}

				if right {
					releaseSegments(segments)
					return true
				}
			}
		}

		_, step := utf8.DecodeRuneInString(s[offset+index:])
		offset += index + step

		releaseSegments(segments)
	}

	return false
}

func (bt BTree) offsetLimit(inputLen int) (offset int, limit int) {
	// bt.Length, bt.RLen and bt.LLen are values meaning the length of runes for each part
	// here we manipulating byte length for better optimizations
	// but these checks still works, cause minLen of 1-rune string is 1 byte.
	if bt.LengthRunes != -1 && bt.LengthRunes > inputLen {
		return 0, 0
	}
	if bt.LeftLengthRunes >= 0 {
		offset = bt.LeftLengthRunes
	}
	if bt.RightLengthRunes >= 0 {
		limit = inputLen - bt.RightLengthRunes
	} else {
		limit = inputLen
	}
	return offset, limit
}

func (bt BTree) String() string {
	const n string = "<nil>"
	var l, r string
	if bt.Left == nil {
		l = n
	} else {
		l = bt.Left.String()
	}
	if bt.Right == nil {
		r = n
	} else {
		r = bt.Right.String()
	}

	return fmt.Sprintf("<btree:[%s<-%s->%s]>", l, bt.Value, r)
}
