package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Printf("r1 area: %v\n", r1.area())
	fmt.Printf("r2 area: %v\n", r2.area())
	fmt.Printf("c1 area: %v\n", c1.area())
	fmt.Printf("c2 area: %v\n", c2.area())
}
