package notify

import (
	"crypto/tls"
	"fmt"
	//"log"
	"tw.com.wd/push/apns/common"
	//"tw.com.wd/push/apns/notification"
	//"github.com/sideshow/apns2"
	//"github.com/sideshow/apns2/certificate"
	"golang.org/x/net/http2"
	"net/http"
	//"io/ioutil"
)

func Push(token, topic, payloadJson string, tlsCert tls.Certificate) {
	fmt.Printf("Token: %v\n", token)
	fmt.Printf("Topic: %v\n", topic)
	fmt.Printf("Payload JSON:%v\n", payloadJson)

	http2Transport := buildHttp2Transport(tlsCert)

	httpClient := http.Client{
		Transport: http2Transport,
	}

	req, err := buildRequest("POST", common.APNS_PRODUCTION_SERVER)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Req: %v\n", req)
	}
}

func buildRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	return req, err
}

func buildTLSClientConfig(tlsCert tls.Certificate) *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
}

func buildHttp2Transport(tlsCert tls.Certificate) *http2.Transport {
	tlsConfig := buildTLSClientConfig(tlsCert)

	return &http2.Transport{
		TLSClientConfig: tlsConfig,
	}
}
