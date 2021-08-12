package _select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer is a convenient function to provide pre-configure timeout racer
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer is the generic racer function with customized timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select will enable you to wait on multiple channels when receive one of the channel the select complete
	select {
	// below 2 cases will determine which url has fastest respond time
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	//	when timeout expires this case will trigger to prevent forever blocking
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	// use go allow concurrent execution
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
