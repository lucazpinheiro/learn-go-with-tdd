// example of DI using the io.Writer interface to pass
// a http.ResponseWriter where data will be written.

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }

// func main() {
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
// }
