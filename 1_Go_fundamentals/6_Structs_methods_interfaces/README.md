# Structs methods and interfaces

## Difference between functions and methods

1. Functions exist in the package.
2. Methods exist in type.
3. One package can not have two same named function.
4. Different type can have same named method but different implementation.

## How to choose between functions and methods

If you need function name to be the same among different types but their implementation is different go for methods.

## Naming convention of the receiver

First letter of the type.

## Type and interface

Using interface we can group related type together, for example:

```
checkArea := func(t testing.TB, shape Shape, want float64) {
    t.Helper()
    got := shape.Area()
    if got != want {
        t.Errorf("got %g want %g", got, want)
    }
}
```

By using `Shape` we can accommodate `Circle`, `Rectangle` and `Triangle` instance, because they all implement `Shape`
interface.

## Table driven tests

[Table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) are useful when you want to build a list of
test cases that can be tested in the same manner.

## Some suggestion to improve test readability

1. Named field like:

```
{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
{shape: Circle{Radius: 10}, want: 314.1592653589793},
{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
```

2. Demonstrate input data not only `got` and `want`

```
%#v got %g want %g
```

3. Use `t.Run` to name test cases.

```
func TestAreaTestDrivenTest(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
```