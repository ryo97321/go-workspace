package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.4)
	fmt.Println("value is", v.Float())
}
