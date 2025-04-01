package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"go_utilities_toolkit/utils/httpclient"
)

func TestSendPostJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	resp, err := httpclient.SendPostJSON(ts.URL, map[string]string{"hello": "world"})
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected result: %v, status: %d", err, resp.StatusCode)
	}
}