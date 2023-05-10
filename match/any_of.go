package match

import "fmt"

type AnyOf struct {
	Matchers Matchers
}

func NewAnyOf(m ...Matcher) AnyOf {
	return AnyOf{Matchers(m)}
}

func (a *AnyOf) Add(m Matcher) error {
	a.Matchers = append(a.Matchers, m)
	return nil
}

func (a AnyOf) Match(s string) bool {
	for _, m := range a.Matchers {
		if m.Match(s) {
			return true
		}
	}

	return false
}

func (a AnyOf) Index(s string) (int, []int) {
	index := -1

	segments := acquireSegments(len(s))
	for _, m := range a.Matchers {
		idx, seg := m.Index(s)
		if idx == -1 {
			continue
		}

		if index == -1 || idx < index {
			index = idx
			segments = append(segments[:0], seg...)
			continue
		}

		if idx > index {
			continue
		}

		// here idx == index
		segments = appendMerge(segments, seg)
	}

	if index == -1 {
		releaseSegments(segments)
		return -1, nil
	}

	return index, segments
}

func (a AnyOf) Len() (l int) {
	l = -1
	for _, m := range a.Matchers {
		ml := m.Len()
		switch {
		case l == -1:
			l = ml
			continue

		case ml == -1:
			return -1

		case l != ml:
			return -1
		}
	}

	return
}

func (a AnyOf) String() string {
	return fmt.Sprintf("<any_of:[%s]>", a.Matchers)
}
