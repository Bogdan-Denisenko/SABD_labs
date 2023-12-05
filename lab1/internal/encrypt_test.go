package internal_test

import (
	"SABD_labs/lab1/internal"
	"io/ioutil"
	"os"
	"testing"
)

func TestXOREncryptDecrypt(t *testing.T) {
	input := []byte("Hello, World!")
	key := byte(0x42)

	encrypted := internal.XorEncryptDecrypt(input, key)

	decrypted := internal.XorEncryptDecrypt(encrypted, key)

	if string(decrypted) != string(input) {
		t.Errorf("XOR encryption/decryption failed. Expected: %s, Got: %s", string(input), string(decrypted))
	}
}

func TestObfuscateAndDeobfuscateXML(t *testing.T) {
	xmlData := []byte(`<root><element>data</element></root>`)
	key := byte(0x42)

	obfuscatedXML, err := obfuscateXMLHelper(xmlData, key)
	if err != nil {
		t.Errorf("Error obfuscating XML: %v", err)
	}

	deobfuscatedXML, err := deobfuscateXMLHelper(obfuscatedXML, key)
	if err != nil {
		t.Errorf("Error deobfuscating XML: %v", err)
	}

	if string(deobfuscatedXML) != string(xmlData) {
		t.Errorf("Obfuscation and deobfuscation failed. Expected: %s, Got: %s", string(xmlData), string(deobfuscatedXML))
	}
}

func obfuscateXMLHelper(xmlData []byte, key byte) ([]byte, error) {
	tmpFile, err := ioutil.TempFile("", "test_obfuscate_*.xml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = ioutil.WriteFile(tmpFile.Name(), xmlData, 0644)
	if err != nil {
		return nil, err
	}

	return internal.ObfuscateXML(tmpFile.Name(), key)
}

func deobfuscateXMLHelper(xmlData []byte, key byte) ([]byte, error) {
	tmpFile, err := ioutil.TempFile("", "test_deobfuscate_*.xml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = ioutil.WriteFile(tmpFile.Name(), xmlData, 0644)
	if err != nil {
		return nil, err
	}

	return internal.DeobfuscateXML(tmpFile.Name(), key)
}
