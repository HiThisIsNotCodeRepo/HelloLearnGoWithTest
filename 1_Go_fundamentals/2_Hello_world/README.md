# Hello, World

## Typical Hello, World

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

## How to test Hello, World

1. Separate "domain" code from "side effect".

> "domain" is the data exists in RAM.
> "side effect" is data exists IO/Network/Screen etc.

```
package hello

// Hello is the data exists in domain
func Hello() string {
	return "Hello, world"
}

```

```
package main

import (
	"fmt"

	"github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/2_Hello_world/hello"
)

func main() {
	fmt.Println(hello.Hello())
}

```

2. Write test

```
package test

import (
	"testing"

	"github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/2_Hello_world/hello"
)

func TestHello(t *testing.T) {
	got := hello.Hello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

```

3. Run test `go test` in test file folder or `go test ./...` in project root folder.

## How to write test

1. File name needs to be `xxx_test.go`.
2. Function name starts with `Test`.
3. Function take one argument `t *testing.T`.
4. Import `"testing"` package.

## Placeholder strings

[Placeholder strings reference](https://pkg.go.dev/fmt#hdr-Printing)

## Refactoring code

1. Add constant to boost performance because it saves creating "Hello, " every time `Hello` is called.

```
package hello

const englishHelloPrefix = "Hello, "

// Hello is the data exists in domain
func Hello(name string) string {
	return englishHelloPrefix + name
}

```

2. Add subtests, using subtests allows us to extract the common code like below.

```
package test

import (
	"testing"

	"github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/2_Hello_world/hello"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello.Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello.Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
```

> Using `testing.TB` allows us to call helper functions from a test or a benchmark.
> `t.Helper()` is needed so when testing fails the line number reported will be our function call rather than tester helper.