package kit

import "syscall/js"

// Env -
var Env *Environment

// Environment -
type Environment struct {
	G *Global
}

func init() {
	Env = &Environment{
		G: NewGlobal(js.Global()),
	}
}

// Wrapper -
type Wrapper interface {
	Value() js.Value
}

// Object -
type Object struct {
	Wrapper
	v js.Value
}

// Value -
func (o *Object) Value() (v js.Value) {
	v = o.v
	return
}

// Set -
func (o *Object) Set(name string, v interface{}) {
	o.Value().Set(name, v)
	return
}

// Get -
func (o *Object) Get(name string) (v js.Value) {
	v = o.Value().Get(name)
	return
}
