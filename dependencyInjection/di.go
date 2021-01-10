package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Ilya")
	http.ListenAndServe(":5000", http.HandlerFunc(GreeterHandler))
}