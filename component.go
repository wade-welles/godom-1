package godom

import (
	"io"
)

// Component .
type Component struct {
	node *Elem
	Quit chan int
}

// HTTP .
func (c *Component) HTTP(method string, url string, body io.Reader) *HTTPRequest {
	r := new(HTTPRequest)
	r.component = c
	r.method = method
	r.url = url
	r.body = body
	return r
}

// WS .
func (c *Component) WS(url string) *WebSocket {
	conn := global.Get("WebSocket").New(url)
	ws := &WebSocket{
		conn: conn,
	}
	return ws
}

func (c *Component) init() {
	c.Quit = make(chan int)
}

func (c *Component) unmount() {
	c.Quit <- 0
	close(c.Quit)
	c.Quit = nil
}
