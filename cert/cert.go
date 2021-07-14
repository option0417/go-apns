package cert

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

func ReadP12FromFile(filepath string, pw string) (tls.Certificate, error) {
	p12ContentBytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)
		return tls.Certificate{}, err
	}

	priKey, cert, err := pkcs12.Decode(p12ContentBytes, pw)

	if err != nil {
		fmt.Println(err)
		return tls.Certificate{}, err
	}

	//fmt.Printf("Subject: %s\n", cert.Subject)
	//fmt.Printf("Issuer: %v\n", cert.Issuer)
	//fmt.Printf("PriKey: %s\n", priKey)

	return tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  priKey,
		Leaf:        cert,
	}, nil
}
