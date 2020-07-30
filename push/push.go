package push

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"tw.com.wd/push/apns/cert"
	"tw.com.wd/push/apns/common"
)

type PushClientBuilder struct {
	pushClient *PushClient
}

func (pcb *PushClientBuilder) Tokens(tokens []string) *PushClientBuilder {
	pcb.pushClient.tokens = tokens
	return pcb
}

func (pcb *PushClientBuilder) Payload(payload string) *PushClientBuilder {
	pcb.pushClient.payload = payload
	return pcb
}

func (pcb *PushClientBuilder) Production() *PushClientBuilder {
	pcb.pushClient.host = common.APNS_DEVELOPMENT_SERVER
	return pcb
}

func (pcb *PushClientBuilder) Build() *PushClient {
	// Fetch cert
	tlsCert, err := cert.ReadP12FromFile(common.CERT_PATH, common.CERT_CODE)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{tlsCert}}
	fmt.Printf("TLS Config:%v\n", tlsConfig)

	// Build http.Client

	return pcb.pushClient
}

type PushClient struct {
	tokens     []string
	payload    string
	host       string
	httpClient *http.Client
}

func (pc *PushClient) Push() {
	fmt.Printf("Do Push\n")
}

func BuildPushClient() *PushClientBuilder {
	return &PushClientBuilder{&PushClient{}}
}
