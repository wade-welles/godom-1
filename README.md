# Godom (Experimental)

Use Godom (not for production) to create user interfaces that run in the browser.

## Documentation

For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/godom).

## Getting Started

```go
package main

import (
	"github.com/twharmon/godom"
)

func main() {
	clicker := newClicker()
	godom.Mount(clicker, godom.Root("#root"))
	<-clicker.Quit
}

type clicker struct {
	godom.Component
	clicked int
}

func newClicker() *clicker {
	return &clicker{}
}

func (c *clicker) Render(root *godom.Elem) {
	p := godom.Create("p").Text(c.clicked)
	btn := godom.Create("button").Text("increment")
	root.AppendElem(p, btn)

	ch := make(chan int)
	btn.OnClick(func(e *godom.MouseEvent) {
		c.clicked++
		ch <- c.clicked
	})

	go func() {
		for {
			select {
			case <-ch:
				p.Text(c.clicked)
			case <-c.Quit:
				return
			}
		}
	}()
}
```

## Contribute

Make a pull request.
