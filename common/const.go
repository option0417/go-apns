package common

//import "fmt"

const (
	APNS_DEVELOPMENT_SERVER = "https://api.sandbox.push.apple.com"
	APNS_PRODUCTION_SERVER  = "https://api.push.apple.com"

	TOPIC = "com.mitake.mitakeeim"

	CERT_PATH = "apns_test.p12"
	CERT_CODE = "86136982"

	Token_OK   = "26c2f5e9f47a7ef0c7fc0d8f178a19712baeda8da11e498026480ed56d219a23"
	Token_Fail = "05599d8ea99c2c924c639c13873393d79abbf29453db952933ed11adeba0b8d8"
	Token_VOIP = "da02c395db49b9918706b34a7dfba7cb670120b6d254be41572c1f960498848a"

	PAYLOAD_A = "{\"aps\":{\"alert\":\"Hello\"}}"
)
