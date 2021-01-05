package main

import "fmt"

type Stringer interface {
	String() string
}

type Human struct {
	Name string
	Age  int
}

type MyString string

type MyInt int

func (h Human) String() string {
	return fmt.Sprintf("%v: %v", h.Name, h.Age)
}

func (ms MyString) String() string {
	return "mystring"
}

func (mi MyInt) String() string {
	return "myint"
}

func F(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	var h Human
	h.Name = "mori"
	h.Age = 24

	var ms MyString = "mystring"
	var mi MyInt = 100

	var s Stringer = h
	F(s)

	s = ms
	F(s)

	s = mi
	F(s)

	fmt.Println()

	F(h)
	F(ms)
	F(mi)

}
