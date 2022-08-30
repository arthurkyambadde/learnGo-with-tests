package main

import (
	"fmt"
)

const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "holla, "
const frenchHelloPrefix = "Bonjour, "

func hello(name string, language string) string {

	if name == "" {
		name = "World!"
	}

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchHelloPrefix + name
	// }

	// return englishHelloPrefix + name
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix

	default:
		prefix = englishHelloPrefix

	}
	return
}

func main() {
	fmt.Println(hello("", ""))

}
