package match

type Super struct{}

func NewSuper() Super {
	return Super{}
}

func (sup Super) Match(s string) bool {
	return true
}

func (sup Super) Len() int {
	return lenNo
}

func (sup Super) Index(s string) (int, []int) {
	segments := acquireSegments(len(s) + 1)
	for i := range s {
		segments = append(segments, i)
	}
	segments = append(segments, len(s))

	return 0, segments
}

func (sup Super) String() string {
	return "<super>"
}
