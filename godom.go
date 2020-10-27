package godom

import (
	"fmt"
	"sync"
	"syscall/js"
	"time"
)

var global = js.Global()
var window = global.Get("window")
var screen = global.Get("screen")
var document = global.Get("document")
var root *Elem
var store = make(map[string][]js.Value)
var storeMu sync.Mutex

var txtStore []js.Value
var txtStoreMu sync.Mutex

func init() {
	t := time.NewTicker(time.Second * 5)
	go func() {
		for {
			<-t.C
			LogGroup("element store counts")
			for k := range store {
				LogDebug(k, "elements:", len(store[k]))
			}
			LogDebug("text elements:", len(txtStore))
			LogGroupEnd()
		}
	}()
}

// Root .
func Root(selector string) *Elem {
	val := document.Call("querySelector", selector)
	if !val.Truthy() {
		return nil
	}
	root = &Elem{
		val: val,
	}
	return root
}

// Mount .
func Mount(u Renderer, e *Elem) {
	e.Clear()
	e.component = u
	u.init()
	u.Render(e)
}

// StaticComponent .
func StaticComponent(elem *Elem) Renderer {
	return &ElemComponent{elem: elem}
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
