package main

import "fmt"

const spanish = "sp"
const russian = "ru"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const russianHelloPrefix = "Привет, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case russian:
		prefix = russianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}