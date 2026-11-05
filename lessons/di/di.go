package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Greet writes a greeting for name to writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler writes an HTTP greeting response.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	handler := http.HandlerFunc(MyGreeterHandler)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
