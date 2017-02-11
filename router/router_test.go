package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestINdexHandler checks that root handler retrns 200
func TestIndexHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(index))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Error Getting Index: %s", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got: %v", res.StatusCode)
	}
}
