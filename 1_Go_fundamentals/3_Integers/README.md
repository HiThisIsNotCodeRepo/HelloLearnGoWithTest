# Integer

## Enough code to make test pass

In the strictest sense of TDD write minimal amount of code to make test pass.

```
func Add(x, y int) int {
	return 4
}
```

## Refactor

```
// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
```

## Add example to test file

```
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//	Output: 6
}
```

> If remove `// Output: 6` `ExampleAdd()` won't execute.