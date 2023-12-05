package internal

import "os"

func XorEncryptDecrypt(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key
	}
	return result
}

func ObfuscateXML(inputFile string, key byte) ([]byte, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	outputXML := XorEncryptDecrypt(xmlData, key)

	return outputXML, nil
}

func DeobfuscateXML(inputFile string, key byte) ([]byte, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	outputXML := XorEncryptDecrypt(xmlData, key)

	return outputXML, nil
}
