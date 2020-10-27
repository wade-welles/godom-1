package godom

import "syscall/js"

// WindowElem .
type WindowElem struct {
	val              js.Value
	popStateListener js.Func
	resizeListener   js.Func
}

// OnPopState .
func (w *WindowElem) OnPopState(f func()) {
	w.popStateListener = js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
		go f()
		return nil
	})
	w.val.Call("addEventListener", "popstate", w.popStateListener)
}

// ClearPopStateListener .
func (w *WindowElem) ClearPopStateListener() {
	w.val.Call("removeEventListener", "popstate", w.popStateListener)
}

// OnResize .
func (w *WindowElem) OnResize(f func()) {
	w.resizeListener = js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
		go f()
		return nil
	})
	w.val.Call("addEventListener", "resize", w.resizeListener)
}

// ClearResizeListener .
func (w *WindowElem) ClearResizeListener() {
	w.val.Call("removeEventListener", "resize", w.resizeListener)
}

// Location .
func (w *WindowElem) Location() *Location {
	return &Location{
		val: w.val.Get("location"),
	}
}

// InnerWidth .
func (w *WindowElem) InnerWidth() int {
	return w.val.Get("innerWidth").Int()
}

// InnerHeight .
func (w *WindowElem) InnerHeight() int {
	return w.val.Get("innerHeight").Int()
}
