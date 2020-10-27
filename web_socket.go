package godom

import (
	"syscall/js"
)

// WebSocket .
type WebSocket struct {
	component *Component
	conn      js.Value
	open      chan int
}

// OnOpen .
func (ws *WebSocket) OnOpen(f func()) {
	ws.conn.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}))
}

// Send .
func (ws *WebSocket) Send(msg string) {
	ws.conn.Call("send", msg)
}

// OnMessage .
func (ws *WebSocket) OnMessage(f func(string)) {
	ws.conn.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0].Get("data").String())
		return nil
	}))
}

// Close .
func (ws *WebSocket) Close() {
	ws.conn.Call("close")
}
