package kit

import "syscall/js"

// Global -
type Global struct {
	Object
}

// NewGlobal -
func NewGlobal(v js.Value) (g *Global) {
	g = new(Global)
	g.v = v
	return
}

// AddEventListener -
func (g *Global) AddEventListener(eventType string, handler interface{}) {
	g.Value().Call("addEventListener", eventType, handler)
}

// Document -
func (g *Global) Document() (d *Document) {
	d = NewDocument(g.Value().Get("document"))
	return
}

// Console -
func (g *Global) Console() (c *Console) {
	c = NewConsole(g.Value().Get("console"))
	return
}

// Console -
type Console struct {
	Object
}

// NewConsole -
func NewConsole(v js.Value) (c *Console) {
	c = new(Console)
	c.v = v
	return
}

// Log -
func (c *Console) Log(msg string) {
	c.Value().Call("log", msg)
}
