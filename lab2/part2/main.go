package main

import (
	"SABD_labs/lab2/file"
	"SABD_labs/lab2/part2/util"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKeyFile := "C:\\Users\\mi\\GolandProjects\\SABD_labs\\lab2\\private_key.pem"
	publicKeyFile := "C:\\Users\\mi\\GolandProjects\\SABD_labs\\lab2\\public_key.pem"

	var privateKey *rsa.PrivateKey
	var publicKey *rsa.PublicKey

	privateKey, _ = file.ReadPrivateKeyFromFile(privateKeyFile)
	publicKey, _ = file.ReadPublicKeyFromFile(publicKeyFile)

	data := []byte("Hello, World!")

	signature, err := util.Sign(privateKey, data)
	if err != nil {
		panic(err)
	}

	verified := util.Verify(publicKey, data, signature)
	fmt.Printf("Data Verified: %t\n", verified)
}
