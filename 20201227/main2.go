package main

import (
	"crypto/rand"
	"fmt"
)

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf)
}

func main() {
	fmt.Println(Key())
}
