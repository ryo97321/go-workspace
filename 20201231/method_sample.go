package main

type T int

func (t *T) f() { println("hi") }

func main() {
	var v T
	(&v).f()
	v.f()
}
