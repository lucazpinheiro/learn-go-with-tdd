package di

import (
	"fmt"
	"io"
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
