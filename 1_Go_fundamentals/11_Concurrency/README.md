# Currency

## Scenario

`CheckWebsites` is used to check the status of a list of URLs

```
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

```

In `results`, `true` represents good response,`false` represents bad response.

## Task 1 mock test

Using real http call to test is time-consuming, instead we should mock.

```
func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}
```

Test function

```
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
```

## Task 2 use benchmark to test the speed of `CheckWebsite`

```
package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

```

To run benchmark using `go test -bench=.` , on Windows Powershell `go test -bench="."`

## Task 3 use concurrency to improve speed

```
package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		// go routine enable concurrent
		go func(u string) {
			// channel avoid race condition, map won't allow concurrent write
			resultChannel <- result{u, wc(u)}
			//	the argument ensure every go routine has its own copy of the url and prevent reuse of url
		}(url)
	}
	// ensure we leave the function before we collect all results
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}

```

Use command to spot race condition `go test -race`