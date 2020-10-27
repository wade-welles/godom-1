package godom

// Renderer .
type Renderer interface {
	Render(*Elem)
	init()
	unmount()
}
