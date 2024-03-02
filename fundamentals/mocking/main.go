/*** MOCKING CHAPTER ***/
package main

import (
	"os"
	"time"
)

// func Countdown(out io.Writer) {
// 	fmt.Fprint(out, "3")
// }

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
