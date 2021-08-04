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

func (pc *PushClient) Push() []*PushResult {
	fmt.Printf("Do Push\n")

	if pc.authToken != "" {
		//c.setTokenHeader(req)
		fmt.Printf("AuthToken: %v\n", pc.authToken)
	}

	reqs := buildRequests(pc)

	if pc.httpClient == nil {
		pc.buildHttpClient()
	}

	var pushResults []*PushResult = make([]*PushResult, len(reqs))

	for i := 0; i < len(reqs); i++ {
		resp, err := pc.httpClient.Do(reqs[i])
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		fmt.Printf("%v\n", resp)

		respContent, err := ioutil.ReadAll(resp.Body)

		pushResults[i] = &PushResult{resp.StatusCode, respContent, resp.Header.Get("Apns-Id")}
		resp.Body.Close()
	}

	return pushResults

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

func buildRequests(pc *PushClient) []*http.Request {
	payloadJson, err := json.Marshal(pc.payload.GetContent())
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	} else {
		fmt.Printf("JSON: %v\n\n", string(payloadJson))
	}

	var reqs []*http.Request = make([]*http.Request, len(pc.tokens))

	for i := 0; i < len(pc.tokens); i++ {
		url := fmt.Sprintf("%v/3/device/%v", pc.host, pc.tokens[i])
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJson))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}

		setupHeaders(req, pc)
		fmt.Printf("Headers: %v\n\n", req.Header)
		fmt.Printf("Req: %v\n\n", req)

		reqs[i] = req
	}

	return reqs
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

// PushResult warp http.Respnose for convenient
type PushResult struct {
	respCode    int
	respContent []byte
	apnsId      string
}

// Methods for PushResult
func (pushResult *PushResult) IsSuccess() bool {
	return pushResult.respCode == 200
}

func (pushResult *PushResult) GetContent() string {
	return string(pushResult.respContent)
}

func (pushResult *PushResult) GetApnsId() string {
	return pushResult.apnsId
}
