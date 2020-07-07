package kit

import "syscall/js"

// Document -
type Document struct {
	Object
}

// NewDocument -
func NewDocument(v js.Value) (d *Document) {
	d = new(Document)
	d.v = v
	return
}

// CreateElement -
func (d *Document) CreateElement(tag string) (e *Element) {
	e = NewElement(d.Value().Call("createElement", tag))
	return
}

// Body -
func (d *Document) Body() (e *Element) {
	e = NewElement(d.Value().Get("body"))
	return
}

// GetElementByID -
func (d *Document) GetElementByID(id string) (e *Element) {
	e = NewElement(d.Value().Call("getElementById", id))
	return
}
