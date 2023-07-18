package main

import "fmt"

const englishHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	var prefix string
	switch language {
	case "Spanish":
		prefix = spanishHello
	case "French":
		prefix = frenchHello
	default:
		prefix = englishHello
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
