package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return fmt.Sprintf("%s%s", prefix, name)
}

// this function was created as the book/tutorial suggests, because the Hello() function
// was getting to big, personally, I don't think this function was necessary, but I will
// write this to follow along the tutorial
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("lucas", "English"))
}
