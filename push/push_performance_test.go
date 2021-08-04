package push

import (
	"testing"
	"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
)

var bp = payload.NewPayloadBuilder().SetAlert("Testing APNs by Go").Build()

func BenchmarkSendNotification(b *testing.B) {

	pushClient := FetchPushClient()

	pushClient.
		Tokens([]string{common.TOKEN_OK}).
		Production().
		PushType(PushTypeAlert).
		Topic(common.TOPIC).
		Payload(bp)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pushClient.Push()
	}
}
