package common

//import "fmt"

const (
	APNS_DEVELOPMENT_SERVER = "https://api.sandbox.push.apple.com"
	APNS_PRODUCTION_SERVER  = "https://api.push.apple.com"

	TOPIC      = "tw.com.mitake.mitakeeim"
	TOPIC_VOIP = "tw.com.mitake.mitakeeim.voip"

	CERT_PATH = "apns_test.p12"
	CERT_CODE = "86136982"

	Token_OK   = "37ff4e443603b62b1a1ece9e078efc07b48a7ef64477bd4d861265cb18d9702c"
	Token_Fail = "05599d8ea99c2c924c639c13873393d79abbf29453db952933ed11adeba0b8d8"
	Token_VOIP = "a5cf04393900483c7b4e40d92afc6ed4db31bbedab774a28877c3631bb4ac562"

	PAYLOAD_A = "{\"aps\":{\"alert\":\"Hello\"}}"
)
