# Array and slices

## Array

Array has a fixed capacity, and we can initialize an array in two ways:

1. `[N]teype{value1,value2,...,valueN}`,`numbers := [5]{1,2,3,4,5}`
2. `[...]type{value1,value2,...,valueN}`,`numbers := [...]{1,2,3,4,5}`

## Range

`range` lets us iterate over array, each iteration `range` returns two values- index and value.

## Arrays and type

Array length is part of their property. If pass an `[4]int` into function expects `[5]int` it won't compile.

## Difference between slice and array

Slices do not encode size and it can have any size.

## How to conduct converage test

`go test -conver`

## Syntax to write variadic function

```
func SumAll(numbersToSum ...[]int) (sums []int)
```

## How to compare slices

Use `reflect.DeepEqual`

```
t.Run("make the sums of some slices", func(t *testing.T) {
    got := SumAllTails([]int{1, 2}, []int{0, 9})
    want := []int{2, 9}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
})
```

But to take note `reflect.DeepEqual` is note "type save"

## `make` key word

`make` allows you to create a slice with starting capacity of `len`