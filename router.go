package godom

import (
	"fmt"
	"regexp"
	"strings"
)

// Route .
type Route struct {
	// Path       string
	Renderer   RouteRenderer
	regexp     *regexp.Regexp
	paramNames []string
}

// Router .
type Router struct {
	routes           []*Route
	notFoundRenderer func() Renderer
	route            string
	newRoute         chan string
	Quit             chan int
}

type param struct {
	key   string
	value string
}

// RouteParams .
type RouteParams []param

// RouteRenderer .
type RouteRenderer func(ps RouteParams) Renderer

// Get .
func (ps RouteParams) Get(key string) string {
	for i := range ps {
		if ps[i].key == key {
			return ps[i].value
		}
	}
	return ""
}

// Mount .
func (r *Router) Mount(outlet *Elem) {
	window := Window()
	window.OnPopState(func() {
		hash := window.Location().Hash()
		hash = strings.Replace(hash, "#", "", 1)
		r.newRoute <- hash
	})

	go func() {
		for {
			select {
			case nextRoute := <-r.newRoute:
				if nextRoute == r.route {
					break
				}
				r.route = nextRoute
				outlet.Clear()
				matched := false
				for _, route := range r.routes {
					if route.regexp.MatchString(r.route) {
						matches := route.regexp.FindAllStringSubmatch(r.route, -1)
						var ps RouteParams
						for i := 1; i < len(matches[0]); i++ {
							ps = append(ps, param{
								key:   route.paramNames[i-1],
								value: matches[0][i],
							})
						}
						Mount(route.Renderer(ps), outlet)
						matched = true
						break
					}
				}
				if !matched && r.notFoundRenderer != nil {
					Mount(r.notFoundRenderer(), outlet) // go ?
				}
			case <-r.Quit:
				window.ClearPopStateListener()
				return
			}
		}
	}()

	location := window.val.Get("location")
	hash := location.Get("hash").String()
	if hash == "" {
		location.Set("href", "/#/")
	} else {
		location.Set("href", location.Get("href"))
	}
}

// NotFound .
func (r *Router) NotFound(renderer func() Renderer) {
	r.notFoundRenderer = renderer
}

var paramNameRegExp = regexp.MustCompile(`{([a-zA-Z0-9-]+):?(.*?)}`)

// NewRouter .
func NewRouter() *Router {
	return &Router{
		routes:   nil,
		newRoute: make(chan string),
		Quit:     make(chan int),
	}
}

// On .
func (r *Router) On(path string, renderer RouteRenderer) {
	rt := new(Route)
	if path[0] != '/' {
		panic("path '" + path + "' does not start with '/'")
	}
	pathRegExpStr := "^" + path + "$"
	matches := paramNameRegExp.FindAllStringSubmatch(path, -1)
	for _, match := range matches {
		if match[2] == "" {
			match[2] = `[^\/]*`
		}
		rt.paramNames = append(rt.paramNames, match[1])
		paramInfoRegExp := regexp.MustCompile(fmt.Sprintf("{%s:?(.*?)}", match[1]))
		pathRegExpStr = paramInfoRegExp.ReplaceAllString(pathRegExpStr, "("+match[2]+")")
	}
	rt.regexp = regexp.MustCompile(pathRegExpStr)
	rt.Renderer = renderer
	r.routes = append(r.routes, rt)
}
