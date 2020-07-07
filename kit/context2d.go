package kit

import "syscall/js"

// Context2d -
type Context2d struct {
	Object
}

// NewContext2d -
func NewContext2d(v js.Value) (c *Context2d) {
	c = &Context2d{}
	c.v = v
	return
}

// FillRect -
func (c *Context2d) FillRect(x, y, w, h int) {
	c.Value().Call("fillRect", x, y, w, h)
}

// ClearRect -
func (c *Context2d) ClearRect(x, y, w, h int) {
	c.Value().Call("clearRect", x, y, w, h)
}

// StrokeRect -
func (c *Context2d) StrokeRect(x, y, w, h int) {
	c.Value().Call("strokeRect", x, y, w, h)
}

// Scale -
func (c *Context2d) Scale(sw float64, sh float64) {
	c.Value().Call("scale", sw, sh)
}
