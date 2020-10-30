package godom

// Component .
type Component interface {
	Render() *Elem
	init()
	unmount()
}
