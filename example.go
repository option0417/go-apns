package main

import (
	"fmt"
	"github.com/sideshow/apns2"
	"log"
	"tw.com.wd/push/apns/cert"
)

func main() {

	var filePath = "apns_test.p12"
	tlsCert, err := cert.ReadP12FromFile(filePath, "86136982")

	var token = "0b02a92dc35d2bd297988b75ead46c356ff8aceb589ee032773eff85d9c1a952"
	var topic = "com.mitake.mitakeeim"
	var payload = "{\"aps\":{\"alert\":\"Hello\"}}"

	notification := &apns2.Notification{}
	notification.DeviceToken = token
	notification.Topic = topic
	notification.Payload = payload

	client := apns2.NewClient(tlsCert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("StatusCode: %v\nAPNS ID: %v\nReason: %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
