package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) GET(pattern string, handle HandlerFunc) {
	e.router.addRoute("GET", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandlerFunc) {
	e.router.addRoute("POST", pattern, handle)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	e.router.handle(context)
}
