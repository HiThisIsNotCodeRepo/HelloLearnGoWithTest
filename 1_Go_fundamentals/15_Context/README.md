# Context

## What is context

> The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.

## How to use context

> At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.

In short, it means pass `context` to every dependent function and use `error` to detect signal that `context` is
cancelled.

## Snippet

```
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprint(w, data)
	}
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			// cancellation branch
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			//	slow append result branch
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	// cancellation branch
	case <-ctx.Done():
		return "", ctx.Err()
	//	slow append result branch
	case res := <-data:
		return res, nil
	}
}
```