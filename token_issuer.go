package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TokenIssuer struct {
	Target        string
	ClientID      string
	clientSecret  string
	SkipSslVerify bool
	CaCertFile    string
	httpClient    *http.Client
}

func NewTokenIssuer(target, clientID, clientSecret string, skipSslValidation bool, caCertFile string) TokenIssuer {
	return TokenIssuer{
		Target:        target,
		ClientID:      clientID,
		clientSecret:  clientSecret,
		SkipSslVerify: skipSslValidation,
		CaCertFile:    caCertFile,
		httpClient:    NewHTTPClient(NewTLSConfig(skipSslValidation, caCertFile)),
	}
}

func (ti TokenIssuer) ClientCredentialsGrant(scopes []string) (*TokenResponse, error) {
	params := url.Values{
		"grant_type": {"client_credentials"},
	}
	return ti.RequestToken(params, scopes)
}

func (ti TokenIssuer) PasswordGrant(username, password string, scopes []string) (*TokenResponse, error) {
	params := url.Values{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
	}
	return ti.RequestToken(params, scopes)
}

func (ti TokenIssuer) RequestToken(params url.Values, scopes []string) (*TokenResponse, error) {
	tokenURL := ti.Target + "/oauth/token"

	if len(scopes) > 0 {
		params["scope"] = scopes
	}

	request, err := http.NewRequest("POST", tokenURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	request.SetBasicAuth(ti.ClientID, ti.clientSecret)

	response, err := ti.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("Invalid status response: %d", response.StatusCode)
	}
	if response.StatusCode == http.StatusNoContent || !strings.Contains(response.Header.Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("Invalid response content type.")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body.")
	}

	tr := NewTokenResponse()
	if err = json.Unmarshal(body, tr); err != nil {
		return nil, fmt.Errorf("Invalid response content.")
	}

	return tr, nil
}
