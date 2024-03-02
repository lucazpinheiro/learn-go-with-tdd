package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	//select allows you to wait on multiple channels. The
	//first one to send a value "wins" and the code underneath
	//the case is executed
	select {
	case <-ping(a):
		fmt.Println("a")
		return a, nil
	case <-ping(b):
		fmt.Println("b")
		return b, nil
	case <-time.After(timeout):
		fmt.Println("oooh")
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	/*
		Notice how we have to use make when creating a channel; rather than say var
		ch chan struct{}. When you use var the variable will be initialised with the
		"zero" value of the type. So for string it is "", int it is 0, etc.

		For channels the zero value is nil and if you try and send to it with <- it
		will block forever because you cannot send to nil channels
	*/
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// first version
// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
