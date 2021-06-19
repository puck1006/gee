package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "the path is %s", req.URL.Path)

	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%s] = %s\n", k, v)
		}
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":9999", &Engine{}))

}
