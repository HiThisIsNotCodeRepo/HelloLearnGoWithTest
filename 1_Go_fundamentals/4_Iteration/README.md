# Iteration

## How to conduct Benchmarking for test

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
```

Key points:

1. To conduct benchmark you need `b *testing.B`.
2. The framework will determine `b.N` for you.
3. To run benchmark use `go test -bench=.` or in Windows Powershell `go test -bench="."`

## Interpret Benchmarking result

```
BenchmarkRepeat-8        8034434               153.1 ns/op
```

`153.1 ns/op` means function takes average 136 nanoseconds to run, and it runs 8034434 times.