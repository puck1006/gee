package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
	fmt.Println("测试git是否推送成功!!!")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the path is %s", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s] = %s\n", k, v)
	}

}
