# Error types

## Use custom error type instead of error string will help

1. Error handling.
2. Easy understand error type and message.

## How to create custom error type

```
type BadStatusError struct {
	URL    string
	Status int
}

func (b BadStatusError) Error() string {
	return fmt.Sprintf("did not get 200 from %s, got %d", b.URL, b.Status)
}
```

## How to convert error to customize type

```
got, isStatusErr := err.(BadStatusError)
```

```
var got BadStatusError
isStatusErr := errors.As(err, &got)
```