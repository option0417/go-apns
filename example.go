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

	var token = "37ff4e443603b62b1a1ece9e078efc07b48a7ef64477bd4d861265cb18d9702c"
	var topic = "tw.com.mitake.mitakeeim"
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
