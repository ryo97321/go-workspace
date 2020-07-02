package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The answer is", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The answer is", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The answer is", m["Answer"])

	if v, ok := m["Answer"]; ok {
		fmt.Println("The answer is", v)
	} else {
		fmt.Println("Not found")
	}
	// fmt.Println("The answer is", v, "Present?", ok)
}
