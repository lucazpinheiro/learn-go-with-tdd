package main

/*** DEPENDENCY INJECTION CHAPTER ***/
// example of DI using the io.Writer interface to pass
// a http.ResponseWriter where data will be written.

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// second version writer has the interface io.Writer
// where any type that implements this interface is valid
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// first version writer is the concrete type bytes.Buffer
// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
