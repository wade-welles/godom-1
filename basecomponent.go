package godom

import (
	"io"
)

// BaseComponent .
type BaseComponent struct {
	node *Elem
	Quit chan int
}

// HTTP .
func (c *BaseComponent) HTTP(method string, url string, body io.Reader) *HTTPRequest {
	r := new(HTTPRequest)
	r.component = c
	r.method = method
	r.url = url
	r.body = body
	return r
}

// WS .
func (c *BaseComponent) WS(url string) *WebSocket {
	conn := global.Get("WebSocket").New(url)
	ws := &WebSocket{
		conn: conn,
	}
	return ws
}

func (c *BaseComponent) init() {
	c.Quit = make(chan int)
}

func (c *BaseComponent) unmount() {
	c.Quit <- 0
	close(c.Quit)
	c.Quit = nil
}
