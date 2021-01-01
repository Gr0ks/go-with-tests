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

	prefix := englishHelloPrefix

	switch language {
	case russian:
		prefix = russianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}