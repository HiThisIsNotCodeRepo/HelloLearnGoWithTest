package poker

import "testing"

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin ant %d", len(store.winCalls), 1)
	}
	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
