package lib

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func NewHTTPClient(tlsConfig *tls.Config) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
			Proxy:           http.ProxyFromEnvironment,
		},
	}
}

func NewTLSConfig(skipSslValidation bool, caCertFile string) *tls.Config {
	config := &tls.Config{
		InsecureSkipVerify: skipSslValidation,
	}
	if !skipSslValidation && caCertFile != "" {
		caCert, err := ioutil.ReadFile(caCertFile)
		if err != nil {
			panic(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		config.RootCAs = caCertPool
	}
	return config
}
