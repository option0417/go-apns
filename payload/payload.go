package payload

type PayloadBuilder struct {
	payload *Payload
}

type Payload struct {
	content map[string]interface{}
}

// Apple-defined keys
type Aps struct {
	Alert            *Alert `json:"alert,omitempty"`
	Badge            int    `json:"badge,omitempty"`
	Sound            *Sound `json:"sound,omitempty"`
	ThreadId         string `json:"thread-id,omitempty"`
	Category         string `json:"category,omitempty"`
	ContentAvailable int    `json:"content-available,omitempty"`
	MutableContent   int    `json:"mutable-content,omitempty"`
	TargetContentId  string `json:"target-content-id,omitempty"`
}

type Alert struct {
	Title           string   `json:"title,omitempty"`
	SubTitle        string   `json:"subtitle,omitempty"`
	Body            string   `json:"body,omitempty"`
	LaunchImage     string   `json:"launch-image,omitempty"`
	TitleLocKey     string   `json:"title-loc-key,omitempty"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`
	SubTitleLocKey  string   `json:"subtitle-loc-key,omitempty"`
	SubTitleLocArgs []string `json:"subtitle-loc-args,omitempty"`
	LocKey          string   `json:"loc-key,omitempty"`
	LocArgs         []string `json:"loc-args,omitempty"`
}

type Sound struct {
	Critical int     `json:"critical,omitempty"`
	Name     string  `json:"name,omitempty"`
	Volume   float32 `json:"volume,omitempty"`
}

// Methods for PayloadBuilder
func BuildPayload() *PayloadBuilder {
	return &PayloadBuilder{
		&Payload{
			map[string]interface{}{
				"aps": &Aps{Alert: &Alert{}, Sound: &Sound{}},
			},
		},
	}
}

func (pb *PayloadBuilder) Build() *Payload {
	return pb.payload
}

// Methods for Payload
func (p *Payload) GetAps() *Aps {
	return p.content["aps"]
}

func (p *Payload) GetAlert() *Alert {
	return p.GetAps().Alert
}

func (p *Payload) GetSound() *Sound {
	return p.GetAps().Sound
}

// Method for Aps
func (pb *PayloadBuilder) Badge(count int) *PayloadBuilder {
	pb.payload.GetAps().Badge = count
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
func (pb *PayloadBuilder) SetAlertTitle(title string) *PayloadBuilder {
	pb.payload.Aps.Alert.Title = title
	return pb
}

func (pb *PayloadBuilder) SetAlertSubTitle(subTitle string) *PayloadBuilder {
	pb.payload.Aps.Alert.SubTitle = subTitle
	return pb
}

func (pb *PayloadBuilder) SetAlertBody(body string) *PayloadBuilder {
	pb.payload.Aps.Alert.Body = body
	return pb
}

func (pb *PayloadBuilder) SetAlertLaunchImageName(launchImage string) *PayloadBuilder {
	pb.payload.Aps.Alert.LaunchImage = launchImage
	return pb
}

func (pb *PayloadBuilder) SetAlertTitleLocKey(titleLocKey string) *PayloadBuilder {
	pb.payload.Aps.Alert.TitleLocKey = titleLocKey
	return pb
}

func (pb *PayloadBuilder) SetAlertTitleLocArgs(titleLocArgs []string) *PayloadBuilder {
	pb.payload.Aps.Alert.TitleLocArgs = titleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetAlertSubTitleLocKey(subTitleKey string) *PayloadBuilder {
	pb.payload.Aps.Alert.SubTitleLocKey = subTitleKey
	return pb
}

func (pb *PayloadBuilder) SetAlertSubTitleLocArgs(subTitleLocArgs []string) *PayloadBuilder {
	pb.payload.Aps.Alert.SubTitleLocArgs = subTitleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetAlertLocKey(locKey string) *PayloadBuilder {
	pb.payload.Aps.Alert.LocKey = locKey
	return pb
}

func (pb *PayloadBuilder) SetAlertLocArgs(locArgs []string) *PayloadBuilder {
	pb.payload.Aps.Alert.LocArgs = locArgs
	return pb
}

// Method for Sound
func (pb *PayloadBuilder) IsCritical(isCritical bool) *PayloadBuilder {
	if isCritical {
		pb.payload.Aps.Sound.Critical = 1
	} else {
		pb.payload.Aps.Sound.Critical = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetSoundName(soundName string) *PayloadBuilder {
	pb.payload.Aps.Sound.Name = soundName
	return pb
}

func (pb *PayloadBuilder) SetVolume(volume float32) *PayloadBuilder {
	pb.payload.Aps.Sound.Volume = volume
	return pb
}

// Method for Custom Property
