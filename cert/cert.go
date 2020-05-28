package cert

import (
	//	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

func ReadP12FromFile(filepath string, pw string) {
	p12Content, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)
		return
	}

	priKey, cert, err := pkcs12.Decode(p12Content, pw)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Subject: %s\n", cert.Subject)
	fmt.Printf("Issuer: %v\n", cert.Issuer)
	fmt.Printf("PriKey: %s\n", priKey)

	/*
		tlsCert := tls.Certificate{
			Certificate: [][]byte{cert.Raw},
			PrivateKey:  priKey,
			Leaf:        cert,
		}
	*/
}
