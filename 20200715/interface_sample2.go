package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// Stringer
func (h Human) String() string {
	return "(" + h.name + " - " + strconv.Itoa(h.age) + " years - " + h.phone + ")"
}

func main() {
	mike := Human{"mike", 25, "222-444-XXX"}
	fmt.Println("This Human is : ", mike)
}
