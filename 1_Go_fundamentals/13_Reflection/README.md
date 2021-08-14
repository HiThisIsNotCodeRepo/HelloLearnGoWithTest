# Reflection

## Scenario

Write a function called `walk(x interface{}, fn func(string)` and calls `fn` to extract all the strings. `x` could be:

1. Basic primitive type like `string`,`int`,`float` etc
2. `Ptr`
3. `Struct`
4. `Slice`
5. `Array`
6. `Map`
7. `Chan`
8. `Func`

## Step 1

### Understand what package we need to use

Obviously when we call this function we wouldn't know what kind of data will be used, therefore we have to use
reflection.

And there are two challenges we need to overcome:

1. Identify the data underlying type.
2. If it's complex type like `Struct`,`Slice` etc we need to do recursive call.

## Step 2

### Coding

```go
package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

```