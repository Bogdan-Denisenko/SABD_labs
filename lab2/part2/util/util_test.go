package util_test

import (
	"SABD_labs/lab2/part2/util"
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := &privateKey.PublicKey

	data := []byte("Hello, world!")

	signature, err := util.Sign(privateKey, data)
	if err != nil {
		t.Fatalf("Signing error: %v", err)
	}

	isValid := util.Verify(publicKey, data, signature)
	if !isValid {
		t.Fatal("Verification failed, expected valid signature")
	}
}

func TestVerifyWithInvalidSignature(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	data := []byte("Hello, world!")

	signature, _ := util.Sign(privateKey, data)

	signature[0] = signature[0] ^ 0xFF

	isValid := util.Verify(publicKey, data, signature)
	if isValid {
		t.Fatal("Verification passed, expected invalid signature")
	}
}

func TestVerifyWithInvalidData(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	data := []byte("Hello, world!")

	signature, _ := util.Sign(privateKey, data)

	modifiedData := append([]byte("Modified "), data...)

	isValid := util.Verify(publicKey, modifiedData, signature)
	if isValid {
		t.Fatal("Verification passed, expected invalid data")
	}
}

