package main

import (
	"bufio"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("Plain Text: ")
	stdin.Scan()
	input := stdin.Text()
	plainText := []byte(input)

	// len(key) -> 8
	key := []byte("12345678")

	block, err := des.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	encryptText := make([]byte, des.BlockSize+len(plainText))
	iv := encryptText[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(encryptText[des.BlockSize:], plainText)
	fmt.Printf("Encrypt: %x\n", encryptText[des.BlockSize:])

	decryptText := make([]byte, len(plainText))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(decryptText, encryptText[des.BlockSize:])
	fmt.Printf("Decrypt: %s\n", string(decryptText))
}
