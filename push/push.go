package push

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
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

type PushResponse struct {
	httpResp *http.Response
}

// Methods for PushClient
func FetchPushClient() *PushClient {
	return &PushClient{}
}

func (pc *PushClient) Tokens(tokens []string) *PushClient {
	pc.tokens = tokens
	return pc
}

func (pc *PushClient) Topic(topic string) *PushClient {
	pc.topic = topic
	return pc
}

func (pc *PushClient) PushType(pushType PushType) *PushClient {
	pc.pushType = pushType
	return pc
}

func (pc *PushClient) Payload(payload *payload.Payload) *PushClient {
	pc.payload = payload
	return pc
}

func (pc *PushClient) Production() *PushClient {
	pc.host = common.APNS_PRODUCTION_SERVER
	return pc
}

func (pc *PushClient) Development() *PushClient {
	pc.host = common.APNS_DEVELOPMENT_SERVER
	return pc
}

func (pc *PushClient) Push() *PushResponse {
	fmt.Printf("Do Push\n")

	payloadJson, err := json.Marshal(pc.payload.GetContent())
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}

	fmt.Printf("JSON: %v\n\n", string(payloadJson))
	fmt.Printf("Token: %v\n\n", pc.tokens[0])

	url := fmt.Sprintf("%v/3/device/%v", pc.host, pc.tokens[0])
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJson))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}

	if pc.authToken != "" {
		//c.setTokenHeader(req)
		fmt.Printf("AuthToken: %v\n", pc.authToken)
	}

	setupHeaders(req, pc)
	fmt.Printf("Headers: %v\n\n", req.Header)
	fmt.Printf("Req: %v\n\n", req)

	pc.buildHttpClient()
	resp, err := pc.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}

	return &PushResponse{resp}

	/*
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
	*/
}

func setupHeaders(req *http.Request, pc *PushClient) {
	req.Header.Set("Content-Type", "applicatiob/json; charset=utf-8")
	req.Header.Set("apns-topic", pc.topic)

	if pc.pushType != "" {
		req.Header.Set("apns-push-type", string(pc.pushType))
	} else {
		req.Header.Set("apns-push-type", string(PushTypeAlert))
	}
}

func (pc *PushClient) buildHttpClient() {
	// Fetch cert
	tlsCert, err := cert.ReadP12FromFile(common.CERT_PATH, common.CERT_CODE)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return
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
	pc.httpClient = httpClient
}

// Methods for PushResponse
func (pushResp *PushResponse) IsSuccess() bool {
	return pushResp.httpResp.StatusCode == 200
}

func (pushResp *PushResponse) GetContent() string {
	defer pushResp.httpResp.Body.Close()
	respContent, err := ioutil.ReadAll(pushResp.httpResp.Body)

	if err != nil {
		fmt.Printf("Read content error, since %v\n", err)
		return ""
	}

	return string(respContent)
}
