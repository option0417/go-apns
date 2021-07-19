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
	Sound            *Sound `json:"sound,omitempty"`
	Badge            int    `json:"badge,omitempty"`
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
func NewPayloadBuilder() *PayloadBuilder {
	return &PayloadBuilder{
		&Payload{
			map[string]interface{}{"aps": &Aps{Alert: &Alert{}, Sound: &Sound{}}},
		},
	}
}

func (pb *PayloadBuilder) Build() *Payload {
	return pb.payload
}

// Methods for Payload
func (p *Payload) GetAps() *Aps {
	return p.content["aps"].(*Aps)
}

func (p *Payload) GetAlert() *Alert {
	return p.GetAps().Alert
}

func (p *Payload) GetSound() *Sound {
	return p.GetAps().Sound
}

func (p *Payload) GetContent() map[string]interface{} {
	return p.content
}

// Method for PayloadBuilder
// Method for Aps from PayloadBuilder
func (pb *PayloadBuilder) Badge(count int) *PayloadBuilder {
	pb.payload.GetAps().Badge = count
	return pb
}

func (pb *PayloadBuilder) SetThreadId(threadId string) *PayloadBuilder {
	pb.payload.GetAps().ThreadId = threadId
	return pb
}

func (pb *PayloadBuilder) SetCategory(category string) *PayloadBuilder {
	pb.payload.GetAps().Category = category
	return pb
}

func (pb *PayloadBuilder) IsContentAvailable(contentAvailable bool) *PayloadBuilder {
	if contentAvailable {
		pb.payload.GetAps().ContentAvailable = 1
	} else {
		pb.payload.GetAps().ContentAvailable = 0
	}
	return pb
}

func (pb *PayloadBuilder) IsMutableContent(mutableContent bool) *PayloadBuilder {
	if mutableContent {
		pb.payload.GetAps().MutableContent = 1
	} else {
		pb.payload.GetAps().MutableContent = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetTargetContentId(targetContentId string) *PayloadBuilder {
	pb.payload.GetAps().TargetContentId = targetContentId
	return pb
}

// Method for Alert from PayloadBuilder
func (pb *PayloadBuilder) SetAlert(alert string) *PayloadBuilder {
	pb.payload.GetAps().Alert.Body = alert
	return pb
}

func (pb *PayloadBuilder) SetAlertTitle(title string) *PayloadBuilder {
	pb.payload.GetAps().Alert.Title = title
	return pb
}

func (pb *PayloadBuilder) SetAlertSubTitle(subTitle string) *PayloadBuilder {
	pb.payload.GetAps().Alert.SubTitle = subTitle
	return pb
}

func (pb *PayloadBuilder) SetAlertBody(body string) *PayloadBuilder {
	pb.payload.GetAps().Alert.Body = body
	return pb
}

func (pb *PayloadBuilder) SetAlertLaunchImageName(launchImage string) *PayloadBuilder {
	pb.payload.GetAps().Alert.LaunchImage = launchImage
	return pb
}

func (pb *PayloadBuilder) SetAlertTitleLocKey(titleLocKey string) *PayloadBuilder {
	pb.payload.GetAps().Alert.TitleLocKey = titleLocKey
	return pb
}

func (pb *PayloadBuilder) SetAlertTitleLocArgs(titleLocArgs []string) *PayloadBuilder {
	pb.payload.GetAps().Alert.TitleLocArgs = titleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetAlertSubTitleLocKey(subTitleKey string) *PayloadBuilder {
	pb.payload.GetAps().Alert.SubTitleLocKey = subTitleKey
	return pb
}

func (pb *PayloadBuilder) SetAlertSubTitleLocArgs(subTitleLocArgs []string) *PayloadBuilder {
	pb.payload.GetAps().Alert.SubTitleLocArgs = subTitleLocArgs
	return pb
}

func (pb *PayloadBuilder) SetAlertLocKey(locKey string) *PayloadBuilder {
	pb.payload.GetAps().Alert.LocKey = locKey
	return pb
}

func (pb *PayloadBuilder) SetAlertLocArgs(locArgs []string) *PayloadBuilder {
	pb.payload.GetAps().Alert.LocArgs = locArgs
	return pb
}

// Method for Sound from PayloadBuilder
func (pb *PayloadBuilder) IsCritical(isCritical bool) *PayloadBuilder {
	if isCritical {
		pb.payload.GetAps().Sound.Critical = 1
	} else {
		pb.payload.GetAps().Sound.Critical = 0
	}
	return pb
}

func (pb *PayloadBuilder) SetSoundName(soundName string) *PayloadBuilder {
	pb.payload.GetAps().Sound.Name = soundName
	return pb
}

func (pb *PayloadBuilder) SetVolume(volume float32) *PayloadBuilder {
	pb.payload.GetAps().Sound.Volume = volume
	return pb
}

// Method for Custom Property from PayloadBuilder
func (pb *PayloadBuilder) SetCustomProperty(key string, val interface{}) *PayloadBuilder {
	pb.payload.content[key] = val
	return pb
}
