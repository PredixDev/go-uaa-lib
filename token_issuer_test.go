package lib_test

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	lib "github.com/PredixDev/go-uaa-lib"
	. "github.com/onsi/gomega"
)

func TestTokenIssuerRequestTokenWithInvalidTarget(t *testing.T) {
	RegisterTestingT(t)

	ti := lib.NewTokenIssuer("Invalid URL", "", "", false, "")
	token, err := ti.RequestToken(url.Values{}, nil)

	Expect(token).To(BeNil())
	Expect(err).ToNot(BeNil())
}

func TestTokenIssuerRequestTokenWithInvalidStatus(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	ti := lib.NewTokenIssuer(server.URL, "", "", false, "")
	token, err := ti.RequestToken(url.Values{}, nil)

	Expect(token).To(BeNil())
	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Invalid status response: 500"))
}

func TestTokenIssuerRequestTokenWithNoContent(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	ti := lib.NewTokenIssuer(server.URL, "", "", false, "")
	token, err := ti.RequestToken(url.Values{}, nil)

	Expect(token).To(BeNil())
	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Invalid response content type."))
}

func TestTokenIssuerRequestTokenWithInvalidContentType(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ti := lib.NewTokenIssuer(server.URL, "", "", false, "")
	token, err := ti.RequestToken(url.Values{}, nil)

	Expect(token).To(BeNil())
	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Invalid response content type."))
}

func TestTokenIssuerRequestTokenWithInvalidBody(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ti := lib.NewTokenIssuer(server.URL, "", "", false, "")
	token, err := ti.RequestToken(url.Values{}, nil)

	Expect(token).To(BeNil())
	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Invalid response content."))
}

func TestTokenIssuerClientCredentialsGrant(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Header.Get("Authorization")).To(Equal("Basic " + base64.StdEncoding.EncodeToString([]byte("client-id:client-secret"))))

		_ = r.ParseForm()
		Expect(r.Form.Get("grant_type")).To(Equal("client_credentials"))
		Expect(r.Form["scope"]).To(ConsistOf("scope1", "scope2"))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token_type": "Token Type", "access_token": "Access Token", "refresh_token": "Refresh Token"}`))
	}))
	defer server.Close()

	ti := lib.NewTokenIssuer(server.URL, "client-id", "client-secret", false, "")
	token, err := ti.ClientCredentialsGrant([]string{"scope1", "scope2"})

	Expect(err).To(BeNil())
	Expect(token).ToNot(BeNil())
	Expect(token.Type).To(Equal("Token Type"))
	Expect(token.Access).To(Equal("Access Token"))
	Expect(token.Refresh).To(Equal("Refresh Token"))
}

func TestTokenIssuerPasswordGrant(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Header.Get("Authorization")).To(Equal("Basic " + base64.StdEncoding.EncodeToString([]byte("client-id:client-secret"))))

		_ = r.ParseForm()
		Expect(r.Form.Get("grant_type")).To(Equal("password"))
		Expect(r.Form.Get("username")).To(Equal("user1"))
		Expect(r.Form.Get("password")).To(Equal("password1"))
		Expect(r.Form["scope"]).To(BeNil())

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token_type": "Token Type", "access_token": "Access Token", "refresh_token": "Refresh Token"}`))
	}))
	defer server.Close()

	ti := lib.NewTokenIssuer(server.URL, "client-id", "client-secret", false, "")
	token, err := ti.PasswordGrant("user1", "password1", nil)

	Expect(err).To(BeNil())
	Expect(token).ToNot(BeNil())
	Expect(token.Type).To(Equal("Token Type"))
	Expect(token.Access).To(Equal("Access Token"))
	Expect(token.Refresh).To(Equal("Refresh Token"))
}
