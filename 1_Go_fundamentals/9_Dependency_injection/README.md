# Dependency Injection

## What can dependency injection can do in testing?

1. Facilitate testing.
2. Write general purpose functions.

## Challenger

Consider testing function `Greet`

```
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

It's difficult to test because `fmt.Printf` prints to stdout, it's difficult for us to capture it. We need to *inject*
dependency of printing. To do that we need function accept an interface rather than concrete type.

## User interface as argument type

```
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

```

We can inject following dependency for `Greet`

1. `bytes.Buffer` to store the string in RAM.
2. `os.Stdout` to print string to stdout.
3. `http.ResponseWriter` to write string to http response.

To summarize dependency injection will allow us to control where these data go.