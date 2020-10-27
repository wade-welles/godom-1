package godom

import "syscall/js"

// KeyboardEvent .
type KeyboardEvent struct {
	val js.Value
}

// Key .
func (e *KeyboardEvent) Key() string {
	return e.val.Get("key").String()
}

// ShiftKey .
func (e *KeyboardEvent) ShiftKey() bool {
	return e.val.Get("shiftKey").Bool()
}

// AltKey .
func (e *KeyboardEvent) AltKey() bool {
	return e.val.Get("altKey").Bool()
}

// CtrlKey .
func (e *KeyboardEvent) CtrlKey() bool {
	return e.val.Get("ctrlKey").Bool()
}

// Repeat .
func (e *KeyboardEvent) Repeat() bool {
	return e.val.Get("repeat").Bool()
}

func newKeyboardEvent(val js.Value) *KeyboardEvent {
	e := new(KeyboardEvent)
	e.val = val
	return e
}
