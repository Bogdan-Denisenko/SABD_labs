package util_test

import (
	"SABD_labs/lab2/part1/util"
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := &privateKey.PublicKey

	plaintext := []byte("Hello, world!")

	ciphertext, err := util.Encrypt(publicKey, plaintext)
	if err != nil {
		t.Fatalf("Encryption error: %v", err)
	}

	decryptedText, err := util.Decrypt(privateKey, ciphertext)
	if err != nil {
		t.Fatalf("Decryption error: %v", err)
	}

	if string(decryptedText) != string(plaintext) {
		t.Fatalf("Decrypted text does not match original plaintext. Expected: %s, Got: %s", plaintext, decryptedText)
	}
}

func TestEncryptAndDecryptWithInvalidKey(t *testing.T) {
	privateKey1, _ := rsa.GenerateKey(rand.Reader, 2048)
	privateKey2, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey1 := &privateKey1.PublicKey

	plaintext := []byte("Hello, world!")

	ciphertext, err := util.Encrypt(publicKey1, plaintext)
	if err != nil {
		t.Fatalf("Encryption error: %v", err)
	}

	_, err = util.Decrypt(privateKey2, ciphertext)
	if err == nil {
		t.Fatal("Expected decryption error with wrong private key, but got nil")
	}
}
