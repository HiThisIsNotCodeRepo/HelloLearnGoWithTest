package revisit_http_handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTeapotHandler(t *testing.T) {
	// test http request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// test http response
	res := httptest.NewRecorder()

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}
