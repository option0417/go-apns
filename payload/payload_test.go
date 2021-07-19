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

func TestSetPayload(t *testing.T) {
	pb := NewPayloadBuilder()

	// Setup string alert
	pb.SetAlert("Hello Go APNS")

	// Setup Alert
	pb.
		SetAlertTitle("Title").
		SetAlertSubTitle("SubTitle").
		SetAlertBody("Body").
		SetAlertLaunchImageName("launch_image").
		SetAlertTitleLocKey("title_loc_key").
		SetAlertTitleLocArgs([]string{"title_loc_arg1", "title_loc_arg2"}).
		SetAlertSubTitleLocKey("sub_title_loc_key").
		SetAlertSubTitleLocArgs([]string{"sub_title_loc_arg1", "sub_title_loc_arg2"}).
		SetAlertLocKey("loc_key").
		SetAlertLocArgs([]string{"loc_arg1", "loc_arg2"})

	// Setup Sound
	pb.
		SetSoundName("sound_name").
		SetVolume(0.5).
		IsCritical(true)

	// Setup Aps
	pb.
		Badge(3).
		SetThreadId("thread_id").
		SetCategory("category").
		IsContentAvailable(true).
		IsMutableContent(true).
		SetTargetContentId("target_content_id")

	// Setup Custom Properties
	pb.
		SetCustomProperty("ptp", 2).
		SetCustomProperty("gi", "G000000800V").
		SetCustomProperty("ci", "T000002X039")

	// Build Payload
	p := pb.Build()

	// Test for Alert
	if p.GetAlert().Title != "Title" {
		t.Errorf("Set Title of Alert fail, except is \"Title\" but got %s", p.GetAlert().Title)
	}

	if p.GetAlert().SubTitle != "SubTitle" {
		t.Errorf("Set SubTitle of Alert fail, except is \"SubTitle\" but got %s", p.GetAlert().SubTitle)
	}

	if p.GetAlert().Body != "Body" {
		t.Errorf("Set Body of Alert fail, except is \"Body\" but got %s", p.GetAlert().Body)
	}

	if p.GetAlert().LaunchImage != "launch_image" {
		t.Errorf("Set LaunchImage of Alert fail, except is \"launch_image\" but got %s", p.GetAlert().LaunchImage)
	}

	if p.GetAlert().TitleLocKey != "title_loc_key" {
		t.Errorf("Set TitleLocKey of Alert fail, except is \"title_loc_key\" but got %s", p.GetAlert().TitleLocKey)
	}

	if p.GetAlert().TitleLocArgs == nil {
		t.Errorf("Set TitleLocArgs of Alert fail, except is not nil.")
	}

	if len(p.GetAlert().TitleLocArgs) != 2 {
		t.Errorf("Set TitleLocArgs of Alert fail, except length is 2 but got %d", len(p.GetAlert().TitleLocArgs))
	}

	if p.GetAlert().SubTitleLocKey != "sub_title_loc_key" {
		t.Errorf("Set SubTitleLocKey of Alert fail, except is \"sub_title_loc_key\" but got %s", p.GetAlert().SubTitleLocKey)
	}

	if p.GetAlert().SubTitleLocArgs == nil {
		t.Errorf("Set SubTitleLocArgs of Alert fail, except is not nil.")
	}

	if len(p.GetAlert().SubTitleLocArgs) != 2 {
		t.Errorf("Set SubTitleLocArgs of Alert fail, except length is 2 but got %d", len(p.GetAlert().SubTitleLocArgs))
	}

	if p.GetAlert().LocKey != "loc_key" {
		t.Errorf("Set LocKey of Alert fail, except is \"loc_key\" but got %s", p.GetAlert().LocKey)
	}

	if p.GetAlert().LocArgs == nil {
		t.Errorf("Set LocArgs of Alert fail, except is not nil.")
	}

	if len(p.GetAlert().LocArgs) != 2 {
		t.Errorf("Set LocArgs of Alert fail, except length is 2 but got %d", len(p.GetAlert().LocArgs))
	}
	//  Test for Sound
	if p.GetSound().Name != "sound_name" {
		t.Errorf("Set name of Sound fail, except is \"sound_name\" but got %s", p.GetSound().Name)
	}

	if p.GetSound().Volume != 0.5 {
		t.Errorf("Set volume of Sound fail, except is \"volume\" but got %f", p.GetSound().Volume)
	}

	if p.GetSound().Critical != 1 {
		t.Errorf("Set critical of Sound fail, except is 1 but got %d", p.GetSound().Critical)
	}

	// Test for Aps
	if p.GetAps().Badge != 3 {
		t.Errorf("Set Badge fail, except is 3 but got %d", p.GetAps().Badge)
	}

	if p.GetAps().ThreadId != "thread_id" {
		t.Errorf("Set ThreadId fail, except is \"thread_id\" but got %s", p.GetAps().ThreadId)
	}

	if p.GetAps().Category != "category" {
		t.Errorf("Set Category fail, except is \"category\" but got %s", p.GetAps().Category)
	}

	if p.GetAps().ContentAvailable != 1 {
		t.Errorf("Set ContentAvailable fail, except is true but got %d", p.GetAps().ContentAvailable)
	}

	if p.GetAps().MutableContent != 1 {
		t.Errorf("Set MutableContent fail, except is true but got %d", p.GetAps().MutableContent)
	}

	if p.GetAps().TargetContentId != "target_content_id" {
		t.Errorf("Set TargetContentId fail, except is \"target_content_id\" but got %s", p.GetAps().TargetContentId)
	}

	// Test for custom-property
	if p.content["ptp"] != 2 {
		t.Errorf("Set custom-property fail, except is 2 but got %d", p.content["ptp"])
	}

	if p.content["gi"] != "G000000800V" {
		t.Errorf("Set custom-property fail, except is \"G000000800V\" but got %s", p.content["gi"])
	}

	if p.content["ci"] != "T000002X039" {
		t.Errorf("Set custom-property fail, except is \"T000002X039\" but got %s", p.content["ci"])
	}

}
