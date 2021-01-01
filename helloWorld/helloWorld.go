package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "sp" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}