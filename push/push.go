package push

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
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

type PushClientBuilder struct {
	pushClient *PushClient
}

// Methods for PushClinetBuilder
func BuildPushClient() *PushClientBuilder {
	return &PushClientBuilder{&PushClient{}}
}

func (pcb *PushClientBuilder) Tokens(tokens []string) *PushClientBuilder {
	pcb.pushClient.tokens = tokens
	return pcb
}

func (pcb *PushClientBuilder) Payload(payload *payload.Payload) *PushClientBuilder {
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

	// Build TLSConfig
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{tlsCert}}

	// Build Transport
	transport := &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Build http.Client
	httpClient := &http.Client{
		Transport: transport,
	}
	pcb.pushClient.httpClient = httpClient

	return pcb.pushClient
}

// Methods for PushClinet
func (pc *PushClient) Push() {
	fmt.Printf("Do Push\n")

	payloadJson, err := json.Marshal(pc.payload)
	if err != nil {
		fmt.printf("Err: %v\n", err)
	}

	url := fmt.Sprintf("%v/3/device/%v", pc.host, n.DeviceToken)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	if c.Token != nil {
		c.setTokenHeader(req)
	}

	setHeaders(req, n)

	httpRes, err := c.requestWithContext(ctx, req)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	response := &Response{}
	response.StatusCode = httpRes.StatusCode
	response.ApnsID = httpRes.Header.Get("apns-id")

	decoder := json.NewDecoder(httpRes.Body)
	if err := decoder.Decode(&response); err != nil && err != io.EOF {
		return &Response{}, err
	}
	return response, nil
}
