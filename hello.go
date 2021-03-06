package main

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// func main() {
// 	// fmt.Println("Hello, world")
// 	// fmt.Println(Hello("world"))
// 	fmt.Println(Hello("world"))
// }

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	// prefix := englishHelloPrefix

	// switch language {
	// case french:
	// 	prefix = frenchHelloPrefix
	// case spanish:
	// 	prefix = spanishHelloPrefix
	// }
	// if language == "Spanish" {
	// 	return spanishHelloPrefix + name
	// }

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
