package kit

import "syscall/js"

// Style -
type Style struct {
	Object
}

// NewStyle -
func NewStyle(v js.Value) (s *Style) {
	s = &Style{}
	s.v = v
	return
}
