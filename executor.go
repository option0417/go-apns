package main

import (
	"fmt"
	//"tw.com.wd/push/apns/cert"
	"tw.com.wd/push/apns/common"
	//"tw.com.wd/push/apns/push"
)

func main() {
	//var filePath = "apns_test.p12"
	//tlsCert, err := cert.ReadP12FromFile(filePath, "86136982")

	/*
		if err != nil {
			fmt.Println("Error: %v\n", err)
			return
		}
	*/

	/*
		var payloadMap = make(map[string]string)
		payloadMap["key1"] = "value1"
		payloadMap["key2"] = "value2"
	*/

	fmt.Printf("Token: %v\n", common.Token_OK)
	fmt.Printf("Topic: %v\n", common.TOPIC)
	fmt.Printf("Payload: %v\n", common.PAYLOAD_A)

	// Build payload
	// payload.PayloadBuilder().tokens([1]string{token}).topic(topic).payload(

	// Build PushClient
	//push.BuildPushClient(token, topic, payload, tlsCert)

	// Do Push

}
