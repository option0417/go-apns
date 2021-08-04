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
		Tokens([]string{common.TOKEN_OK}).
		Production().
		PushType(PushTypeAlert).
		Topic(common.TOPIC).
		Payload(p)

	pushResults := pushClient.Push()

	if pushResults == nil {
		t.Errorf("Send notification failed, since response is nil")
	}

	for i := 0; i < len(pushResults); i++ {
		if !pushResults[i].IsSuccess() {
			t.Errorf("Send notification failed, since %s", pushResults[i].GetContent())
		} else {
			t.Logf("APNS ID: %s", pushResults[i].GetApnsId())
		}
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

	pushResults := pushClient.Push()

	if pushResults == nil {
		t.Errorf("Send notification failed, since response is nil")
	}

	for i := 0; i < len(pushResults); i++ {
		if pushResults[i].IsSuccess() {
			t.Errorf("Send notification should fail since empty device-token")
		} else {
			t.Log(pushResults[i].GetContent())
		}
	}
}

func TestSendMultiNotification(t *testing.T) {
	pushClient := FetchPushClient()

	pushClient.
		Tokens([]string{common.TOKEN_OK, common.TOKEN_OK, common.TOKEN_OK}).
		Production().
		PushType(PushTypeAlert).
		Topic(common.TOPIC).
		Payload(p)

	pushResults := pushClient.Push()

	if pushResults == nil {
		t.Errorf("Send notification failed, since response is nil")
	}

	if len(pushResults) != 3 {
		t.Errorf("Except receive 3 results, but got %d.", len(pushResults))
	}

	for i := 0; i < len(pushResults); i++ {
		if !pushResults[i].IsSuccess() {
			t.Errorf("Send notification failed, since %s", pushResults[i].GetContent())
		} else {
			t.Logf("APNS ID: %s", pushResults[i].GetApnsId())
		}
	}
}

func TestSendMultiNotificationWithWrongToken(t *testing.T) {
	pushClient := FetchPushClient()

	pushClient.
		Tokens([]string{common.TOKEN_OK, "", common.TOKEN_OK}).
		Production().
		PushType(PushTypeAlert).
		Topic(common.TOPIC).
		Payload(p)

	pushResults := pushClient.Push()

	if pushResults == nil {
		t.Errorf("Send notification failed, since response is nil")
	}

	if len(pushResults) != 3 {
		t.Errorf("Except receive 3 results, but got %d.", len(pushResults))
	}

	for i := 0; i < len(pushResults); i++ {
		if !pushResults[i].IsSuccess() && i != 1 {
			t.Errorf("Send notification failed, since %s", pushResults[i].GetContent())
		}
	}
}
