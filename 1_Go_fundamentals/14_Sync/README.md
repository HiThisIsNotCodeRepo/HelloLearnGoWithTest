# Sync

## Useful Sync package

1. `sync.WaitGroup` waits for a collection of goroutines to finish.
2. `sync.WaitGroup.Add` adds delta, which may be negative, to the WaitGroup counter.
3. `sync.WaitGroup.Done` decrements the WaitGroup counter by one.
4. `sync.WaitGroup.Wait` blocks until the WaitGroup counter is zero.
5. `sync.Mutex` is a mutual exclusion lock.

## How does sync.Mutex take effect

When it locks only one goroutine will be allowed to execute. The other goroutines have to wait the previous goroutine
release lock.

## Don't embed `sync.Mutex` into struct

```
type Counter struct {
	sync.Mutex
	value int
}
```

If you do this the `Lock` and `Unlock` will become `Counter` method, and they are supposed to be used internally.

## Don't copy mutex

> A Mutex must not be copied after first use.

Use `go vet` to check any bug.

## When to use mutex or channel

- Pass the value use channel.
- Mutate value use mutex.