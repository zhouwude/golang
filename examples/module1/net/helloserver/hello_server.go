// hello_server.go
package main

import (
	"fmt"
	"net/http"
)

// http.Handler是一个接口 实现了ServeHTTP 方法就行
type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}

// Output in browser-window with url http://localhost:4000:  Hello!
