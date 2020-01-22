package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet greets
func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler is handler for http requests
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, "world")
}

func main() {
    http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}