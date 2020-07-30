package main

import (
	"fmt"
	"tw.com.wd/push/apns/push"
)

func main() {
	// Build payload
	// payload.PayloadBuilder().tokens([1]string{token}).topic(topic).payload(

	// Build PushClient
	pc := push.
		BuildPushClient().
		Tokens([]string{common.Token_OK}).
		Payload(common.PAYLOAD_A).
		Production().
		Build()

	// Do Push
	pc.Push()
}
