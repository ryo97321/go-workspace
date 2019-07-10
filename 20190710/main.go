package main

import "fmt"

type Struct struct {
	V float64
}

func main() {
	var i1, i2, i3 interface{}
	i1 = 12
	i2 = "interface"
	i3 = Struct{V: 3.14}
	fmt.Println(i1, i2, i3)
}
