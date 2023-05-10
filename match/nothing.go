package match

type Nothing struct{}

func NewNothing() Nothing {
	return Nothing{}
}

func (Nothing) Match(s string) bool {
	return len(s) == 0
}

func (Nothing) Index(s string) (int, []int) {
	return 0, segments0
}

func (Nothing) Len() int {
	return lenZero
}

func (Nothing) String() string {
	return "<nothing>"
}
