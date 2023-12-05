package main

import (
	"SABD_labs/lab2/file"
	"SABD_labs/lab2/part1/util"
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

	encryptedData, err := util.Encrypt(publicKey, data)
	if err != nil {
		panic(err)
	}

	decryptedData, err := util.Decrypt(privateKey, encryptedData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original Data: %s\n", data)
	fmt.Printf("Decrypted Data: %s\n", decryptedData)
}
