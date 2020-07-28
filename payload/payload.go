package payload

type Payload struct {
	aps   Aps
	alert Alert
	sound Sound
}

func (p *Payload) Alert(alert interface{}) *Payload {
	return p
}

func (p *Payload) Badge(number int) *Payload {
	return p
}

func (p *Payload) Sound(sound interface{}) *Payload {
	return p
}
