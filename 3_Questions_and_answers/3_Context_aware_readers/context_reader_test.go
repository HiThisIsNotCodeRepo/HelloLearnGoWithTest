package context_aware_reader

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("lets just see how a normal reader works", func(t *testing.T) {
		rdr := strings.NewReader("123456")
		got := make([]byte, 3)
		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}
		assertBufferHas(t, got, "123")
		_, err = rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "456")
	})
	t.Run("behaves like a normal reader", func(t *testing.T) {
		rdr := NewCancellableReader(context.Background(), strings.NewReader("123456"))
		got := make([]byte, 3)
		_, err := rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}
		assertBufferHas(t, got, "123")

		_, err = rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "456")
	})
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
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
