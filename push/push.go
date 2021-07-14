package push

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tw.com.wd/push/apns/cert"
	"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
)

type PushType string

const (
	PushTypeAlert        PushType = "alert"
	PushTypeBackground   PushType = "background"
	PushTypeVoip         PushType = "voip"
	PushTypeComplication PushType = "complication"
	PushTypeFileProvider PushType = "fileprovider"
	PushTypeMdm          PushType = "mdm"
)

type PushClient struct {
	tokens     []string
	payload    *payload.Payload
	host       string
	pushType   PushType
	apnsId     string
	expiration time.Time
	priority   int
	topic      string
	collapseId string
	authToken  string
	httpClient *http.Client
}

// Methods for PushClient
func FetchPushClient() *PushClient {
	return &PushClient{}
}

func (pc *PushClient) Tokens(tokens []string) *PushClient {
	pc.tokens = tokens
	return pc
}

func (pc *PushClient) Payload(payload *payload.Payload) *PushClient {
	pc.payload = payload
	return pc
}

func (pc *PushClient) Production() *PushClient {
	pc.host = common.APNS_DEVELOPMENT_SERVER
	return pc
}

func (pc *PushClient) Push() {
	// Fetch cert
	tlsCert, err := cert.ReadP12FromFile(common.CERT_PATH, common.CERT_CODE)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{tlsCert}}
	fmt.Printf("TLS Config:%v\n", tlsConfig)

	// Do Push
	fmt.Printf("Do Push\n")
	payloadJson, err := json.Marshal(pc.payload)

	if err == nil {
		fmt.Println(string(payloadJson))
	}
}
