package main

import (
	"fmt"
	"tw.com.wd/push/apns/cert"
)

func main() {
	fmt.Println("vim-go")

	var filePath = "apns_test.p12"

	cert.ReadP12FromFile(filePath, "86136982")
}
