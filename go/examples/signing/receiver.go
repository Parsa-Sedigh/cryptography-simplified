package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func receive(signature, digest []byte) {
	publicKeyBytes, err := os.ReadFile("signing/keys/public.pem")
	if err != nil {
		log.Fatalf("Error reading public key: %v", err)
	}

	pemBlock, _ := pem.Decode(publicKeyBytes)
	publicKey, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}

	if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest, signature); err != nil {
		log.Fatalf("Error verifying signature: %v", err)
	}

	log.Println("signature verified. Yay!")
}
