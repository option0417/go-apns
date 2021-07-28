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
		t.Errorf("Send notification failed, since %s", resp.GetContent())
	} else {
		t.Log(resp.GetContent())
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
		t.Errorf("Send notification should fail since empty device-token but get %s", resp.GetContent())
	} else {
		t.Log(resp.GetContent())
	}
}
