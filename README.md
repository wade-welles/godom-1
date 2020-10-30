# Godom (Experimental)

Use Godom (not for production) to create user interfaces that run in the browser.

## Documentation

For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/godom).

## Getting Started

### Basic Usage

```go
package main

import (
	"github.com/twharmon/godom"
)

func main() {
	clicker := newClicker()
	godom.Root("#root").Append(clicker)
	<-clicker.Quit
}

type clicker struct {
	godom.BaseComponent
	clicked int
}

func newClicker() *clicker {
	return &clicker{}
}

func (c *clicker) Render() *godom.Elem {
	p := godom.Create("p").Text(c.clicked)
	btn := godom.Create("button").Text("increment")
    root := godom.Create("div")
    root.Appen(btn, p)

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
    
    return root
}
```

### Template
See [basic template](https://github.com/twharmon/godom-template) for more.

## Contribute

Make a pull request.
