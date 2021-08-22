# Context aware reader

## The benefit to use context aware reader

We can't guarantee how long it will take to finish reading to set it with expiry time, we can make use of `context`.

## A reader with context

```
func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return &readerCtx{
		ctx:      ctx,
		delegate: rdr,
	}
}

type readerCtx struct {
	ctx      context.Context
	delegate io.Reader
}

func (r *readerCtx) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.delegate.Read(p)
}

```

## Test case

```
	t.Run("stops reading when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		// will trigger candle after 1 second pass
		//time.AfterFunc(1*time.Second, cancel)
		// will trigger candle after 4 second fail
		time.AfterFunc(4*time.Second, cancel)
		rdr := NewCancellableReader(ctx, strings.NewReader("123456"))
		got := make([]byte, 3)
		_, err := rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "123")
		time.Sleep(2 * time.Second)
		//cancel()
		n, err := rdr.Read(got)
		if err == nil {
			t.Error("expected an error after cancellation but didn't got one")
		}
		if n > 0 {
			t.Errorf("expected 0 bytes to be read after cancellation but %d were read", n)
		}
	})
```