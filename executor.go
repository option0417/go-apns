package main

import (
	"fmt"
	"tw.com.wd/push/apns/cert"
	"tw.com.wd/push/apns/notification"
)

func main() {
	fmt.Println("vim-go")

	var filePath = "apns_test.p12"

	cert.ReadP12FromFile(filePath, "86136982")

	fmt.Printf("\n")

	fmt.Printf("Token OK: %v\n", notification.Token_OK)
	fmt.Printf("Token Fail: %v\n", notification.Token_Fail)
	fmt.Printf("Token VOIP: %v\n", notification.Token_VOIP)
}
