package main

import (
	"fmt"
	"tw.com.wd/push/apns/cert"
	"tw.com.wd/push/apns/notify"
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

	notify.Push(token, topic, payload, tlsCert)
}
