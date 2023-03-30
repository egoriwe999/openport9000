package main

import (
	"fmt"
	"net/http"
)
func main() {
	fmt.Print("Go project listening in Docker on 9000 port")
		handler := HttpHandler{}
		http.ListenAndServe(":9000", handler)
}

type HttpHandler struct{}
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Created by egoriwe999")
	res.Write(data)
}
