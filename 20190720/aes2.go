package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("Input plain Text: ")
	stdin.Scan()
	input := stdin.Text()

	plainText := []byte(input)

	key := []byte("0123456789ABCDEF0123456789ABCDEF")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	encryptText := make([]byte, aes.BlockSize+len(plainText))

	iv := encryptText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("encryptText[:aes.BlockSize]: %x\n", encryptText[:aes.BlockSize])
	// fmt.Printf("iv: %x\n", iv)

	// Encrypt
	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(encryptText[aes.BlockSize:], plainText)
	fmt.Printf("Encrypt: %x\n", encryptText)

	// Decrypt
	decryptText := make([]byte, len(encryptText[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, encryptText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptText, encryptText[aes.BlockSize:])
	fmt.Printf("Decrypt: %s\n", string(decryptText))
}
