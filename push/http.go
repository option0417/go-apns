package push

import (
	"crypto/tls"
	//"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

const (
	Method_POST   = "POST"
	Method_GET    = "GET"
	Method_PUT    = "PUT"
	Method_DELETE = "DELETE"
)

func FetchHttpClient() *http.Client {
	return &http.Client{}
}

func buildRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	return req, err
}

func buildHttp2Transport(tlsCert tls.Certificate) *http2.Transport {
	tlsConfig := buildTLSClientConfig(tlsCert)

	return &http2.Transport{
		TLSClientConfig: tlsConfig,
	}
}

func buildTLSClientConfig(tlsCert tls.Certificate) *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
}
