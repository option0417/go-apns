package push

import (
	"crypto/tls"
	"fmt"
	"net/http"
	//"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
)

type PushClient struct {
	tokens     []string
	payload    *payload.Payload
	pushUrl    string
	httpClient *http.Client
}

func BuildPushClient(tokens []string, payload *payload.Payload, pushUrl string, tlsCert tls.Certificate) *PushClient {
	fmt.Printf("Token: %v\n", tokens)

	httpClient := FetchHttpClient()

	return &PushClient{tokens, payload, pushUrl, httpClient}
}
