package push

import (
	"testing"
	"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
)

var p = payload.NewPayloadBuilder().SetAlert("Testing APNs by Go").Build()

func TestSendNotification(t *testing.T) {
	pushClient := FetchPushClient()

	pushClient.
		Tokens([]string{common.Token_OK}).
		Production().
		PushType(PushTypeAlert).
		Topic(common.TOPIC).
		Payload(p)

	resp := pushClient.Push()

	if !resp.IsSuccess() {
		t.Errorf("Send notification failed")

	}
}

func TestSendNotificationWithWrongToken(t *testing.T) {
	pushClient := FetchPushClient()

	pushClient.
		Tokens([]string{""}).
		Production().
		PushType(PushTypeAlert).
		Topic(common.TOPIC).
		Payload(p)

	resp := pushClient.Push()

	if resp.IsSuccess() {
		t.Errorf("Send notification should fail since empty device-token")

	}
}
