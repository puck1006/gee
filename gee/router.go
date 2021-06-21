package gee

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handle HandlerFunc) {
	log.Printf("Route %s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handle
}

func (r *router) handle(c *Context) {
	fmt.Println("触发了")
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Printf("404 Not Found %s\n", c.Path)
	}
}
