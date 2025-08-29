package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

var msg = "Hello World!"

func send() ([]byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("error generating RSA private key:", err)
	}

	publicKey := &privateKey.PublicKey

	// save RSA private key as PEM
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	if err := os.WriteFile("signing/keys/private.pem", privateKeyPem, 0644); err != nil {
		log.Fatalf("Error saving private key: %v", err)
	}

	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	if err := os.WriteFile("signing/keys/public.pem", publicKeyPem, 0644); err != nil {
		log.Fatalf("Error saving public key: %v", err)
	}

	hash := sha256.New()

	_, err = hash.Write([]byte(msg))
	if err != nil {
		log.Fatal("error hashing message:", err)
	}

	digest := hash.Sum(nil)

	hash.Reset()

	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, digest)
	if err != nil {
		log.Fatal("error signing message:", err)
	}

	log.Printf("digest: %s\n", string(digest))
	log.Printf("signature: %s\n", string(signature))

	return signature, digest
}
