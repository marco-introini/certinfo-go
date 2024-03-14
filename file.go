package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func fileInfo(filename string) {

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist:", filename)
			os.Exit(1)
		} else {
			fmt.Println("Error checking file:", err)
			os.Exit(1)
		}
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(f)

	pemData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		os.Exit(1)
	}

	// Check if the file is a valid PEM certificate.
	block, rest := pem.Decode(pemData)
	if len(rest) != 0 || block == nil {
		fmt.Println("File is not a valid PEM certificate")
		os.Exit(1)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("File is not a valid PEM certificate:", err)
		os.Exit(1)
	}

	fmt.Println("Common Name:", cert.Subject)
	fmt.Println("Issuer:", cert.Issuer)
	fmt.Println("Expiration:", cert.NotAfter)

}
