package lib_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	lib "github.com/PredixDev/go-uaa-lib"
	. "github.com/onsi/gomega"
)

func TestInfoServerSuccess(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.URL.RequestURI()).To(Equal("/login"))
		Expect(r.Header.Get("Accept")).To(Equal("application/json"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	info := lib.NewInfoClient(server.URL, false, "")

	Expect(info.Server()).To(BeNil())
}

func TestInfoServerError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	info := lib.NewInfoClient(server.URL, false, "")

	err := info.Server()
	Expect(err.Error()).To(Equal("Invalid status response: 500"))
}
