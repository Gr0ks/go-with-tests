package main

import "fmt"

const spanish = "sp"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}