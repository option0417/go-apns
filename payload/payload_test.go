package payload

import (
	"testing"
)

func TestGetPayloadBuilder(t *testing.T) {
	var pb *PayloadBuilder = nil
	pb = NewPayloadBuilder()

	if pb == nil {
		t.Errorf("Instance return from payload.NewPayloadBuilder() is nil")
	}

}

func TestPayloadBuilderContainPayload(t *testing.T) {
	var pb *PayloadBuilder = nil
	pb = NewPayloadBuilder()

	if pb.payload == nil {
		t.Errorf("Payload which in the PayloadBuilder is nil")
		return
	}

	if pb.payload.content == nil {
		t.Errorf("Content of the payload is nil")
		return
	}

	if pb.payload.content["aps"] == nil {
		t.Errorf("Aps of the Content of the payload is nil")
		return
	}

	aps := pb.payload.content["aps"].(*Aps)
	if aps.Alert == nil {
		t.Errorf("Alert of the Aps of the Content of the payload is nil")
		return
	}

	if aps.Sound == nil {
		t.Errorf("Sound of the Aps of the Content of the payload is nil")
		return
	}
}
