package common

//import "fmt"

const (
	APNS_DEVELOPMENT_SERVER = "https://api.sandbox.push.apple.com"
	APNS_PRODUCTION_SERVER  = "https://api.push.apple.com"

	TOPIC = "com.mitake.mitakeeim"

	CERT_PATH = "apns_test.p12"
	CERT_CODE = "86136982"

	Token_OK   = "0b02a92dc35d2bd297988b75ead46c356ff8aceb589ee032773eff85d9c1a952"
	Token_Fail = "05599d8ea99c2c924c639c13873393d79abbf29453db952933ed11adeba0b8d8"
	Token_VOIP = "da02c395db49b9918706b34a7dfba7cb670120b6d254be41572c1f960498848a"

	PAYLOAD_A = "{\"aps\":{\"alert\":\"Hello\"}}"
)
