package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

func main() {
	plainText := []byte("This is a test!!")

	// use AES128 -> len(key) = 16
	// use AES192 -> len(key) = 24
	// use AES256 -> len(key) = 32
	key := []byte("0123456789ABCDEF0123456789ABCDEF")

	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt
	encryptText := make([]byte, len(plainText))
	cipher.Encrypt(encryptText, plainText)
	fmt.Printf("encrypted: %x\n", encryptText)

	// Decrypt
	decryptText := make([]byte, len(encryptText))
	cipher.Decrypt(decryptText, encryptText)
	fmt.Printf("decrypted: %s\n", string(decryptText))
}
