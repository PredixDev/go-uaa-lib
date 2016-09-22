package lib

import (
	"fmt"
	"net/http"
)

type InfoClient interface {
	Server() error
}

type infoClient struct {
	Client     *http.Client
	URL        string
	SkipSsl    bool
	CaCertFile string
}

var InfoClientFactory InfoClientFactoryInterface = infoClientFactory{}

type InfoClientFactoryInterface interface {
	New(url string, skipSslValidation bool, caCertFile string) InfoClient
}

type infoClientFactory struct{}

func (f infoClientFactory) New(url string, skipSslValidation bool, caCertFile string) InfoClient {
	return infoClient{
		Client:     NewHTTPClient(NewTLSConfig(skipSslValidation, caCertFile)),
		URL:        url,
		SkipSsl:    skipSslValidation,
		CaCertFile: caCertFile,
	}
}

func (info infoClient) Server() error {
	req, err := http.NewRequest("GET", info.URL+"/login", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := info.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("Invalid status response: %d", resp.StatusCode)
	}
	_ = resp.Body.Close()
	return nil
}
