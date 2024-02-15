package main

import "fmt"

const (
	french = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

// a public function(starts with uppercase)=>exposed to world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// a private function(starts with lowercase)
func greetingPrefix(language string) (prefix string) {
	switch language {
        case "French":
                prefix = frenchHelloPrefix
        case "Spanish":
                prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // we have a var prefix so no need for "return prefix"
}

func main() {
	fmt.Println(Hello("Munga", "Spanish"))
}
