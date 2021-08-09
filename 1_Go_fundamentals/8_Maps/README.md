# Maps

## How map stores items

The map works in a similar way as dictionary, `key` like word and `value` like definition.

## Key type requirement

The key type is special, it can only be a comparable type. Value type can be any type.

## Map value in method

This can mutate map.

```
func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}
```

Because a map value is the address of map structure.

## Proper way to initialize empty map

```
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)

```

## Create constant error

```
const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

```

## Why Delete doesn't need to return error

Because the value to be deleted no matter exist or not , won't change the final result, means after delete no value
exists.