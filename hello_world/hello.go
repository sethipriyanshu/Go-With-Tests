package main

const englisHelloPrefix = "Hello! "
const spanishHelloPrefix = "Hola! "
const frenchHelloPrefix = "Bonjour! "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englisHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {

}
