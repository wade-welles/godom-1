package godom

import (
	"strings"
	"syscall/js"
)

// Elem .
type Elem struct {
	name              string
	val               js.Value
	children          []*Elem
	components        []Component
	clickListener     js.Func
	inputListener     js.Func
	keydownListener   js.Func
	mousemoveListener js.Func
	attrs             []string
	isTextNode        bool
}

// Clear .
func (e *Elem) Clear() {
	for _, c := range e.components {
		c.unmount()
	}
	e.components = nil
	for _, child := range e.children {
		e.removeChild(child)
		child.Clear()
	}
	e.children = nil
}

// Render .
func (e *Elem) Render() *Elem {
	return e
}

func (e *Elem) appendElem(children ...*Elem) *Elem {
	for _, child := range children {
		e.val.Call("appendChild", child.val)
		e.children = append(e.children, child)
	}
	return e
}

// Text .
func (e *Elem) Text(text interface{}) *Elem {
	if len(e.children) == 1 && e.children[0].isTextNode {
		e.children[0].val.Set("nodeValue", text)
		return e
	}
	e.Clear()
	n := createTextElem(text)
	e.val.Call("appendChild", n.val)
	e.children = append(e.children, n)
	return e
}

// Attr .
func (e *Elem) Attr(name string, value interface{}) *Elem {
	e.val.Set(name, value)
	e.attrs = append(e.attrs, name)
	return e
}

// OnClick .
func (e *Elem) OnClick(cb func(*MouseEvent)) *Elem {
	e.removeListener("click", e.clickListener)
	f := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		go cb(newMouseEvent("click", args[0]))
		return nil
	})
	e.val.Call("addEventListener", "click", f)
	e.clickListener = f
	return e
}

// OnInput .
func (e *Elem) OnInput(cb func(string)) *Elem {
	e.removeListener("input", e.inputListener)
	f := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		go cb(args[0].Get("target").Get("value").String())
		return nil
	})
	e.val.Call("addEventListener", "input", f)
	e.inputListener = f
	return e
}

// OnKeyDown .
func (e *Elem) OnKeyDown(cb func(*KeyboardEvent)) *Elem {
	e.removeListener("keydown", e.keydownListener)
	f := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		go cb(newKeyboardEvent(args[0]))
		return nil
	})
	e.val.Call("addEventListener", "keydown", f)
	e.keydownListener = f
	return e
}

// OnMouseMove .
func (e *Elem) OnMouseMove(cb func(*MouseEvent)) *Elem {
	e.removeListener("mousemove", e.mousemoveListener)
	f := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		go cb(newMouseEvent("mousemove", args[0]))
		return nil
	})
	e.val.Call("addEventListener", "mousemove", f)
	e.mousemoveListener = f
	return e
}

// Style .
func (e *Elem) Style(name string, value string) *Elem {
	e.val.Get("style").Set(name, value)
	e.registerAttr("style")
	return e
}

// Classes .
func (e *Elem) Classes(names ...string) *Elem {
	e.val.Set("classList", strings.Join(names, " "))
	e.registerAttr("class")
	return e
}

// AddClass .
func (e *Elem) AddClass(name string) *Elem {
	e.val.Get("classList").Call("add", name)
	e.registerAttr("class")
	return e
}

// RemoveClass .
func (e *Elem) RemoveClass(name string) *Elem {
	e.val.Get("classList").Call("remove", name)
	return e
}

// ToggleClass .
func (e *Elem) ToggleClass(name string) *Elem {
	e.val.Get("classList").Call("toggle", name)
	e.registerAttr("class")
	return e
}

// Select .
func (e *Elem) Select(selector string) *Elem {
	for _, child := range e.children {
		if e := child.find(selector); e != nil {
			return e
		}
	}
	return nil
}

func (e *Elem) find(selector string) *Elem {
	if e.val.Call("matches", selector).Bool() {
		return e
	}
	for _, child := range e.children {
		if e := child.find(selector); e != nil {
			return e
		}
	}
	return nil
}

func (e *Elem) registerAttr(name string) {
	for _, attr := range e.attrs {
		if attr == name {
			return
		}
	}
	e.attrs = append(e.attrs, name)
}

func (e *Elem) removeChild(child *Elem) {
	child.removeListener("click", child.clickListener)
	child.removeListener("input", child.inputListener)
	child.removeListener("keydown", child.keydownListener)
	child.removeListener("mousemove", child.mousemoveListener)
	for i := 0; i < len(child.attrs); i++ {
		child.val.Call("removeAttribute", child.attrs[i])
	}
	e.val.Call("removeChild", child.val)
	if child.isTextNode {
		txtStoreMu.Lock()
		child.val.Set("nodeValue", nil)
		txtStore = append(txtStore, child.val)
		txtStoreMu.Unlock()
	} else {
		storeMu.Lock()
		store[child.name] = append(store[child.name], child.val)
		storeMu.Unlock()
	}
}

func (e *Elem) removeListener(ty string, f js.Func) {
	if f.Truthy() {
		e.val.Call("removeEventListener", ty, f)
		switch ty {
		case "click":
			e.clickListener = js.FuncOf(nil)
		case "input":
			e.inputListener = js.FuncOf(nil)
		case "mousemove":
			e.mousemoveListener = js.FuncOf(nil)
		case "keydown":
			e.keydownListener = js.FuncOf(nil)
		}
	}
}
