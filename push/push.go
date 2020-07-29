package push

import (
	//"crypto/tls"
	//"fmt"
	"net/http"
	"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
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
	pcb.pushClient.pushUrl = common.APNS_DEVELOPMENT_SERVER
	return pcb
}

func (pcb *PushClientBuilder) Build() *PushClient {
	return pcb.pushClient
}

type PushClient struct {
	tokens     []string
	payload    *payload.Payload
	pushUrl    string
	httpClient *http.Client
}

func BuildPushClient() *PushClientBuilder {
	return &PushClientBuilder{&PushClient{}}
}
