package lib_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	lib "github.com/PredixDev/go-uaa-lib"
)

func TestServerSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RequestURI() != "/login" {
			t.Error("URL: expected /login, found", r.URL)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Error("Accept header: expected application/json, found", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	info := lib.NewInfoClient(server.URL, false, "")

	err := info.Server()
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
}

func TestServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	info := lib.NewInfoClient(server.URL, false, "")

	err := info.Server()
	if err.Error() != "Invalid status response: 500" {
		t.Errorf("Invalid response: expected 'Invalid status response: 500', found '%s'", err.Error())
	}
}
