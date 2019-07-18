package main

import (
	"bytes"
	"fmt"
)

func main() {
	byteString1 := []byte("This is a test")
	byteString2 := []byte("This is a test")
	fmt.Printf("%T %v\n", byteString1, byteString2)

	judge := bytes.Compare(byteString1, byteString2)
	if judge == 0 {
		fmt.Println("same")
	} else {
		fmt.Println("diff")
	}

	if bytes.ContainsAny(byteString1, "This") {
		fmt.Println("byteString1 contains This")
	}

	fmt.Printf("%T %v\n", 'T', 'T')
	if bytes.ContainsRune(byteString1, 'T') {
		fmt.Println("byteString1 contains T")
	}

	isNumber := bytes.Count(byteString1, []byte("is"))
	fmt.Printf("is counts %d\n", isNumber)

	if bytes.Equal(byteString1, byteString2) {
		fmt.Println("same")
	} else {
		fmt.Println("diff")
	}

	hasIs := bytes.HasPrefix(byteString1, []byte("is"))
	if hasIs {
		fmt.Println("byteString1 has is")
	} else {
		fmt.Println("hyteString1 don't have is")
	}

	isPlaceIndex := bytes.Index(byteString1, []byte("is"))
	fmt.Printf("is at %d\n", isPlaceIndex)

}
