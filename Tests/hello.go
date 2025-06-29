package main

import (
	"fmt"
	"strings"
)

const (
	spanish          = "spanish"
	french           = "french"
	spanishToEnglish = "Hola "
	frenchToEnglish  = "Bonjour "
	defaultword      = "Hello "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greeter(strings.ToLower(language)) + name
}

func greeter(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishToEnglish
	case french:
		prefix = frenchToEnglish
	default:
		prefix = defaultword
	}
	return
}
func main() {
	fmt.Println(Hello("World", "English"))
}
