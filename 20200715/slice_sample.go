package main

import "fmt"

func main() {
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice)

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	var a, b []byte

	a = ar[:5]
	b = ar[5:]

	fmt.Printf("ar: %v\n", ar)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	a[0] = 'z'
	fmt.Printf("ar: %v\n", ar)
	fmt.Printf("a: %v\n", a)

	a = append(a, 'X')
	fmt.Printf("a: %v\n", a)
}
