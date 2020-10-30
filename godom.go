package godom

import (
	"fmt"
	"sync"
	"syscall/js"
)

var global = js.Global()
var window = global.Get("window")
var screen = global.Get("screen")
var document = global.Get("document")

var store = make(map[string][]js.Value)
var storeMu sync.Mutex

var txtStore []js.Value
var txtStoreMu sync.Mutex

// Root .
func Root(selector string) *Elem {
	val := document.Call("querySelector", selector)
	if !val.Truthy() {
		return nil
	}
	return &Elem{
		val: val,
	}
}

// Renderer .
type Renderer interface {
	Render() *Elem
}

// Append .
func (e *Elem) Append(renderers ...Renderer) *Elem {
	for _, renderer := range renderers {
		componentRenderer, ok := renderer.(Component)
		if ok {
			e.components = append(e.components, componentRenderer)
			componentRenderer.init()
			e.appendElem(componentRenderer.Render())
		} else {
			e.appendElem(renderer.Render())
		}
	}
	return e
}

// Window .
func Window() *WindowElem {
	return &WindowElem{
		val: window,
	}
}

// Create .
func Create(tag string) *Elem {
	e := new(Elem)
	e.name = tag
	storeMu.Lock()
	defer storeMu.Unlock()
	vals := store[tag]
	if len(vals) > 0 {
		e.val, vals = vals[0], vals[1:]
		store[tag] = vals
	} else {
		e.val = document.Call("createElement", tag)
	}
	return e
}

func createTextElem(text interface{}) *Elem {
	var txtNode js.Value
	txtStoreMu.Lock()
	if len(txtStore) > 0 {
		txtNode, txtStore = txtStore[0], txtStore[1:]
		txtNode.Set("nodeValue", text)
	} else {
		txtNode = document.Call("createTextNode", text)
	}
	txtStoreMu.Unlock()
	return &Elem{
		val:        txtNode,
		isTextNode: true,
	}
}

// RouteTo .
func RouteTo(path string) {
	window.Get("location").Set("href", fmt.Sprintf("/#%s", path))
}
