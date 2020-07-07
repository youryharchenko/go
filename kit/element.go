package kit

import "syscall/js"

// Element -
type Element struct {
	Object
}

// NewElement -
func NewElement(v js.Value) (e *Element) {
	e = new(Element)
	e.v = v
	return
}

// SetInnerHTML -
func (e *Element) SetInnerHTML(v string) {
	e.Value().Set("innerHTML", v)
}

// AppendChild -
func (e *Element) AppendChild(c *Element) {
	e.Value().Call("appendChild", c.Value())
	return
}

// GetContext -
func (e *Element) GetContext(t string) (c *Context2d) {
	c = NewContext2d(e.Value().Call("getContext", t))
	return
}

// ClientWidth -
func (e *Element) ClientWidth() (w int) {
	w = e.Get("clientWidth").Int()
	return
}

// ClientHeight -
func (e *Element) ClientHeight() (h int) {
	h = e.Get("clientHeight").Int()
	return
}

// Style -
func (e *Element) Style() (s *Style) {
	s = NewStyle(e.Value().Get("style"))
	return
}
