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

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == russian {
		return russianHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}