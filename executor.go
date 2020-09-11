package main

import (
	//	"encoding/json"
	//	"fmt"
	"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
	"tw.com.wd/push/apns/push"
)

func main() {
	// Build payload
	pb := payload.BuildPayload()

	// Setup Alert
	pb.
		SetAlertTitle("Title").
		SetAlertSubTitle("SubTitle").
		SetAlertBody("Body")
	/*
		SetLaunchImageName("launch_image").
		SetTitleLocKey("title_loc_key").
		SetTitleLocArgs([]string{"title_loc_arg1", "title_loc_arg2"}).
		SetSubTitleLocKey("sub_title_loc_key").
		SetSubTitleLocArgs([]string{"sub_title_loc_arg1", "sub_title_loc_arg2"}).
		SetLocKey("loc_key").
		SetLocArgs([]string{"loc_arg1", "loc_arg2"})
	*/

	// Setup Sound
	pb.
		SetSoundName("sound_name").
		SetVolume(0.5).
		IsCritical(true)

	// Setup Aps
	pb.
		Badge(3)
		//		SetThreadId("thread_id").
		//		SetCategory("category").
		//		IsContentAvailable(true).
		//		IsMutableContent(true)

	// Setup Custom Properties
	pb.
		SetCustomProperty("ptp", 2).
		SetCustomProperty("gi", "G000000800V").
		SetCustomProperty("ci", "T000002X039")

	p := pb.Build()
	// Build PushClient

	pc := push.
		BuildPushClient().
		Tokens([]string{common.Token_VOIP}).
		Payload(p).
		Production().
		PushType(push.PushTypeVoip).
		Topic("com.mitake.mitakeeim.voip").
		Build()

	// Do Push
	pc.Push()

}
