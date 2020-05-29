package main

import (
	"fmt"
	"log"
	"tw.com.wd/push/apns/cert"
	//"tw.com.wd/push/apns/notification"
	"github.com/sideshow/apns2"
	//"github.com/sideshow/apns2/certificate"
	//"golang.org/x/net/http2"
	//"io/ioutil"
	//"net/http"
)

func main() {

	var filePath = "apns_test.p12"

	tlsCert, err := cert.ReadP12FromFile(filePath, "86136982")

	var token = "47e9b9c8ffb7cc4310f72189ef9d5cc9cf56d1c286083f5e41118c1f1b91c273"
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

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
