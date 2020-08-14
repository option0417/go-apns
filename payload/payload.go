package payload

import (
	"fmt"
)

type PayloadBuilder struct {
	payload *Payload
}

type Payload struct {
	Aps   *Aps
	Alert *Alert
	Sound *Sound
}

// Apple-defined keys
type Aps struct {
	Alert            *Alert
	Badge            int
	Sound            *Sound
	ThreadId         string
	Category         string
	ContentAvailable int
	MutableContent   int
	TargetContentId  string
}

type Alert struct {
	Title           string
	SubTitle        string
	Body            string
	LaunchImage     string
	TitleLocKey     string
	TitleLocArgs    []string
	SubTitleLocKey  string
	SubTitleLocArgs []string
	LocKey          string
	LocArgs         []string
}

type Sound struct {
	Critical int
	Name     string
	Volume   float32
}

// Methods for PayloadBuilder
func BuildPayload() *PayloadBuilder {
	return &PayloadBuilder{&Payload{&Aps{}, &Alert{}, &Sound{}}}
}

func (pb *PayloadBuilder) Build() *Payload {
	pb.payload.Aps.Alert = pb.payload.Alert
	pb.payload.Aps.Sound = pb.payload.Sound

	pb.payload.Alert = nil
	pb.payload.Sound = nil

	return pb.payload
}

// Methods for Payload
func (p *Payload) GetAps() *Aps {
	return p.Aps
}

func (p *Payload) GetAlert() *Alert {
	return p.Alert
}

func (p *Payload) GetSound() *Sound {
	return p.Sound
}

// Method for Aps
func (pb *PayloadBuilder) Badge(count int) *PayloadBuilder {
	pb.payload.Aps.Badge = count
	return pb
}

func (pb *PayloadBuilder) SetThreadId(threadId string) *PayloadBuilder {
	pb.payload.Aps.ThreadId = threadId
	return pb
}

func (pb *PayloadBuilder) SetCategory(category string) *PayloadBuilder {
	pb.payload.Aps.Category = category
	return pb
}

func (pb *PayloadBuilder) IsContentAvailable(contentAvailable bool) *PayloadBuilder {
	if contentAvailable {
		pb.payload.Aps.ContentAvailable = 1
	} else {
		pb.payload.Aps.ContentAvailable = 0
	}
	return pb
}

func (pb *PayloadBuilder) IsMutableContent(mutableContent bool) *PayloadBuilder {
	if mutableContent {
		pb.payload.Aps.MutableContent = 1
	} else {
		pb.payload.Aps.MutableContent = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetTargetContentId(targetContentId string) *PayloadBuilder {
	pb.payload.Aps.TargetContentId = targetContentId
	return pb
}

// Method for Alert
func (pb *PayloadBuilder) SetTitle(title string) *PayloadBuilder {
	pb.payload.Alert.Title = title
	return pb
}

func (pb *PayloadBuilder) SetSubTitle(subTitle string) *PayloadBuilder {
	pb.payload.Alert.SubTitle = subTitle
	return pb
}

func (pb *PayloadBuilder) SetBody(body string) *PayloadBuilder {
	pb.payload.Alert.Body = body
	return pb
}

func (pb *PayloadBuilder) SetLaunchImageName(launchImage string) *PayloadBuilder {
	pb.payload.Alert.LaunchImage = launchImage
	return pb
}

func (pb *PayloadBuilder) SetTitleLocKey(titleLocKey string) *PayloadBuilder {
	pb.payload.Alert.TitleLocKey = titleLocKey
	return pb
}

func (pb *PayloadBuilder) SetTitleLocArgs(titleLocArgs []string) *PayloadBuilder {
	pb.payload.Alert.TitleLocArgs = titleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetSubTitleLocKey(subTitleKey string) *PayloadBuilder {
	pb.payload.Alert.SubTitleLocKey = subTitleKey
	return pb
}

func (pb *PayloadBuilder) SetSubTitleLocArgs(subTitleLocArgs []string) *PayloadBuilder {
	pb.payload.Alert.SubTitleLocArgs = subTitleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetLocKey(locKey string) *PayloadBuilder {
	pb.payload.Alert.LocKey = locKey
	return pb
}

func (pb *PayloadBuilder) SetLocArgs(locArgs []string) *PayloadBuilder {
	pb.payload.Alert.LocArgs = locArgs
	return pb
}

// Nethod for Sound
func (pb *PayloadBuilder) IsCritical(isCritical bool) *PayloadBuilder {
	if isCritical {
		pb.payload.Sound.Critical = 1
	} else {
		pb.payload.Sound.Critical = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetSoundName(soundName string) *PayloadBuilder {
	pb.payload.Sound.Name = soundName
	return pb
}

func (pb *PayloadBuilder) SetVolume(volume float32) *PayloadBuilder {
	pb.payload.Sound.Volume = volume
	return pb
}

func (s *Sound) String() string {
	return fmt.Sprintf("Critical: %v, Name: %v, Volume: %v\n", s.Critical, s.Name, s.Volume)
}
