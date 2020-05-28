package notification

type Token struct {
	val   string
	_type byte
}

const (
	Token_OK   = "47e9b9c8ffb7cc4310f72189ef9d5cc9cf56d1c286083f5e41118c1f1b91c273"
	Token_Fail = "05599d8ea99c2c924c639c13873393d79abbf29453db952933ed11adeba0b8d8"
	Token_VOIP = "da02c395db49b9918706b34a7dfba7cb670120b6d254be41572c1f960498848a"
)
