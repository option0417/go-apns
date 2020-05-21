package payload

type PayloadBuilder struct {
	aps   Aps
	alert Alert
	sound Sound
}

func (pb *PayloadBuilder) Alert(alert interface{}) *PayloadBuilder {
	return pb
}

func (pb *PayloadBuilder) Badge(number int) *PayloadBuilder {
	return pb
}

func (pb *PayloadBuilder) Sound(sound interface{}) *PayloadBuilder {
	return pb
}
