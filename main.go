package main

import (
	"net/http"
)

var version = "undefined"

//HTTPHandler create a handler struct
type HTTPHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data := []byte("Hello, hello, hello, how low - GO! - " + version) // slice of bytes
	// write `data` to response
	res.Write(data)
}

func main() {
	// create a new handler
	handler := HTTPHandler{}
	// listen and serve
	http.ListenAndServe(":3001", handler)
}
