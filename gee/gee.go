package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRoute(method string, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handle
}

func (e *Engine) GET(pattern string, handle HandlerFunc) {
	e.addRoute("GET", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandlerFunc) {
	e.addRoute("POST", pattern, handle)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		for key, value := range e.router {
			fmt.Printf("%s=>%s\n", key, value)
		}
		fmt.Printf("404 Not Found %s\n", req.URL.Path)
	}
}

// key := req.Method + "-" + req.URL.Path
// if handler, ok := engine.router[key]; ok {
// 	handler(w, req)
// } else {
// 	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
// }
