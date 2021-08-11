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
