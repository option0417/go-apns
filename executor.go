package main

import (
	//	"encoding/json"
	//	"fmt"
	//"tw.com.wd/push/apns/common"
	"tw.com.wd/push/apns/payload"
	"tw.com.wd/push/apns/push"
)

func main() {
	// Build payload
	pb := payload.BuildPayload()

	// Setup Alert
	pb.
		//	SetAlertTitle("Title").
		//	SetAlertSubTitle("SubTitle").
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
		SetSoundName("ringtone_006.mp3")
		//SetVolume(0.5).
		//IsCritical(true)

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
		SetCustomProperty("ei", "G000000800V_T000002H04S_E00001p10Dx").
		SetCustomProperty("ti", "T000002H04S").
		SetCustomProperty("v", 1).
		SetCustomProperty("v2", 1).
		SetCustomProperty("tp", "00").
		SetCustomProperty("ptp", 1).
		SetCustomProperty("s2", 1)

	p := pb.Build()
	// Build PushClient

	pc := push.
		BuildPushClient().
		Tokens([]string{"19e8ca26952adbd37328e1317b172f8250ed82d0d77edec3a1f3347ecadc6e1d"}).
		Payload(p).
		Production().
		PushType(push.PushTypeVoip).
		Topic("com.mitake.mitakeeim").
		Build()

	// Do Push
	pc.Push()

}
