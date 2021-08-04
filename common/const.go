package common

const (
	APNS_DEVELOPMENT_SERVER = "https://api.sandbox.push.apple.com"
	APNS_PRODUCTION_SERVER  = "https://api.push.apple.com"

	TOPIC      = "tw.com.mitake.mitakeeim"
	TOPIC_VOIP = "tw.com.mitake.mitakeeim.voip"

	CERT_PATH = "apns_test.p12"
	CERT_CODE = "86136982"

	TOKEN_OK   = "1022c086ffdf047cc499b96c59c8e7eed8b8f274d6d5ec6926a3db8a5d38bc86"
	TOKEN_Fail = "05599d8ea99c2c924c639c13873393d79abbf29453db952933ed11adeba0b8d8"
	TOKEN_VOIP = "a5cf04393900483c7b4e40d92afc6ed4db31bbedab774a28877c3631bb4ac562"

	PAYLOAD_A = "{\"aps\":{\"alert\":\"Hello\"}}"
)
