package main

import (
	"fmt"
	"tw.com.wd/push/apns/cert"
	//"tw.com.wd/push/apns/push"
)

func main() {
	var filePath = "apns_test.p12"
	tlsCert, err := cert.ReadP12FromFile(filePath, "86136982")

	if err != nil {
		fmt.Println("Error: %v\n", err)
		return
	}

	var token = "47e9b9c8ffb7cc4310f72189ef9d5cc9cf56d1c286083f5e41118c1f1b91c273"
	var topic = "com.mitake.mitakeeim"
	var payload = "{\"aps\":{\"alert\":\"Hello\"}}"
	var payloadMap = make(map[string]string)
	payloadMap["key1"] = "value1"
	payloadMap["key2"] = "value2"

	// Build payload
	// payload.PayloadBuilder().tokens([1]string{token}).topic(topic).payload(

	// Build PushClient
	//push.BuildPushClient(token, topic, payload, tlsCert)

	// Do Push

}
