package main

import (
	"SABD_labs/lab1/internal"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "", "Specify 'obfuscate' or 'deobfuscate'")
	inputFile := flag.String("input", "", "Specify the input XML file")
	outputFile := flag.String("output", "", "Specify the output XML file")
	key := flag.Uint("key", 0x42, "Specify the XOR key as an integer")

	flag.Parse()

	if *mode == "" || *inputFile == "" || *outputFile == "" {
		fmt.Println("Usage: go run main.go -mode=[obfuscate|deobfuscate] -input=input.xml -output=output.xml -key=42")
		os.Exit(1)
	}

	switch *mode {
	case "obfuscate":
		obfuscatedXML, err := internal.ObfuscateXML(*inputFile, byte(*key))
		if err != nil {
			fmt.Println("Error obfuscating XML:", err)
			os.Exit(1)
		}

		err = os.WriteFile(*outputFile, obfuscatedXML, 0644)
		if err != nil {
			fmt.Println("Error writing obfuscated XML to file:", err)
			os.Exit(1)
		}

		fmt.Println("Obfuscation complete. Output written to", *outputFile)

	case "deobfuscate":
		deobfuscatedXML, err := internal.DeobfuscateXML(*inputFile, byte(*key))
		if err != nil {
			fmt.Println("Error deobfuscating XML:", err)
			os.Exit(1)
		}

		err = os.WriteFile(*outputFile, deobfuscatedXML, 0644)
		if err != nil {
			fmt.Println("Error writing deobfuscated XML to file:", err)
			os.Exit(1)
		}

		fmt.Println("Deobfuscation complete. Output written to", *outputFile)

	default:
		fmt.Println("Invalid mode. Use 'obfuscate' or 'deobfuscate'.")
		os.Exit(1)
	}
}
