package payload

type PayloadBuilder struct {
	payload *Payload
}

type Payload struct {
	aps   *Aps
	alert *Alert
	sound *Sound
}

// Apple-defined keys
type Aps struct {
	alert            *Alert
	badge            int
	sound            *Sound
	threadId         string
	category         string
	contentAvailable int
	mutableContent   int
	targetContentId  string
}

type Alert struct {
	title           string
	subTitle        string
	body            string
	launchImage     string
	titleLocKey     string
	titleLocArgs    []string
	subTitleLocKey  string
	subTitleLocArgs []string
	locKey          string
	locArgs         []string
}

type Sound struct {
	critical int
	name     string
	volume   float32
}

// Methods for PayloadBuilder
func BuildPayload() *PayloadBuilder {
	return &PayloadBuilder{&Payload{&Aps{}, &Alert{}, &Sound{}}}
}

func (pb *PayloadBuilder) Build() *Payload {
	pb.payload.aps.alert = pb.payload.alert
	pb.payload.aps.sound = pb.payload.sound
	return pb.payload
}

func (pb *PayloadBuilder) Aps() *Aps {
	return pb.payload.aps
}

func (pb *PayloadBuilder) Alert() *Alert {
	return pb.payload.alert
}

func (pb *PayloadBuilder) Sound() *Sound {
	return pb.payload.sound
}

// Method for Aps
func (pb *PayloadBuilder) Badge(count int) *PayloadBuilder {
	pb.payload.aps.badge = count
	return pb
}

func (pb *PayloadBuilder) SetThreadId(threadId string) *PayloadBuilder {
	pb.payload.aps.threadId = threadId
	return pb
}

func (pb *PayloadBuilder) SetCategory(category string) *PayloadBuilder {
	pb.payload.aps.category = category
	return pb
}

func (pb *PayloadBuilder) IsContentAvailable(contentAvailable bool) *PayloadBuilder {
	if contentAvailable {
		pb.payload.aps.contentAvailable = 1
	} else {
		pb.payload.aps.contentAvailable = 0
	}
	return pb
}

func (pb *PayloadBuilder) IsMutableContent(mutableContent bool) *PayloadBuilder {
	if mutableContent {
		pb.payload.aps.mutableContent = 1
	} else {
		pb.payload.aps.mutableContent = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetTargetContentId(targetContentId string) *PayloadBuilder {
	pb.payload.aps.targetContentId = targetContentId
	return pb
}

// Method for Alert
func (pb *PayloadBuilder) SetTitle(title string) *PayloadBuilder {
	pb.payload.alert.title = title
	return pb
}

func (pb *PayloadBuilder) SetSubTitle(subTitle string) *PayloadBuilder {
	pb.payload.alert.subTitle = subTitle
	return pb
}

func (pb *PayloadBuilder) SetBody(body string) *PayloadBuilder {
	pb.payload.alert.body = body
	return pb
}

func (pb *PayloadBuilder) SetLaunchImageName(launchImage string) *PayloadBuilder {
	pb.payload.alert.launchImage = launchImage
	return pb
}

func (pb *PayloadBuilder) SetTitleLocKey(titleLocKey string) *PayloadBuilder {
	pb.payload.alert.titleLocKey = titleLocKey
	return pb
}

func (pb *PayloadBuilder) SetTitleLocArgs(titleLocArgs []string) *PayloadBuilder {
	pb.payload.alert.titleLocArgs = titleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetSubTitleLocKey(subTitleKey string) *PayloadBuilder {
	pb.payload.alert.subTitleLocKey = subTitleKey
	return pb
}

func (pb *PayloadBuilder) SetSubTitleLocArgs(subTitleLocArgs []string) *PayloadBuilder {
	pb.payload.alert.subTitleLocArgs = subTitleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetLocKey(locKey string) *PayloadBuilder {
	pb.payload.alert.locKey = locKey
	return pb
}

func (pb *PayloadBuilder) SetLocArgs(locArgs []string) *PayloadBuilder {
	pb.payload.alert.locArgs = locArgs
	return pb
}

// Nethod for Sound
func (pb *PayloadBuilder) IsCritical(isCritical bool) *PayloadBuilder {
	if isCritical {
		pb.payload.sound.critical = 1
	} else {
		pb.payload.sound.critical = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetSoundName(soundName string) *PayloadBuilder {
	pb.payload.sound.name = soundName
	return pb
}

func (pb *PayloadBuilder) SetVolume(volume float32) *PayloadBuilder {
	pb.payload.sound.volume = volume
	return pb
}
