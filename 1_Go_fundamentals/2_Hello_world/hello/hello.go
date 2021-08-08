package hello

const englishHelloPrefix = "Hello, "

// Hello is the data exists in domain
func Hello(name string) string {
	return englishHelloPrefix + name
}
