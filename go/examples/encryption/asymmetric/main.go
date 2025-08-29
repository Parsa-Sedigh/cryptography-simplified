package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

var msg = "Hello World"

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	if err := os.WriteFile("encryption/asymmetric/keys/private.pem", privateKeyPEM, 0644); err != nil {
		log.Fatalf("Error writing private.pem: %v", err)
	}

	publicKey := &privateKey.PublicKey
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	if err := os.WriteFile("encryption/asymmetric/keys/public.pem", publicKeyPEM, 0644); err != nil {
		log.Fatalf("Error writing public.pem: %v", err)
	}

	cipherMsg, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(msg))
	if err != nil {
		log.Fatalf("Error encrypting message: %v", err)
	}

	fmt.Printf("msg was %s, cipher is: %s:\n", msg, cipherMsg)

	plainMsg, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherMsg)
	if err != nil {
		log.Fatalf("Error decrypting message: %v", err)
	}

	fmt.Printf("cipher was %s,\n msg is: %s\n", cipherMsg, plainMsg)
}
