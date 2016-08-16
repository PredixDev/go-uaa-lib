package lib

import (
	"fmt"
	"net/http"
)

type InfoClient struct {
	Client     *http.Client
	URL        string
	SkipSsl    bool
	CaCertFile string
}

func NewInfoClient(url string, skipSslValidation bool, caCertFile string) InfoClient {
	return InfoClient{
		Client:     NewHTTPClient(NewTLSConfig(skipSslValidation, caCertFile)),
		URL:        url,
		SkipSsl:    skipSslValidation,
		CaCertFile: caCertFile,
	}
}

func (info InfoClient) Server() error {
	req, err := http.NewRequest("GET", info.URL+"/login", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := info.Client.Do(req)
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("Invalid status response: %d", resp.StatusCode)
	}
	if err != nil {
		return err
	}
	_ = resp.Body.Close()
	return nil
}
