package main

import "fmt"

func main() {
	m := map[string]int{"x": 10, "y": 20, "z": 50}

	if n, ok := m["z"]; ok {
		fmt.Printf("m[\"z\"] = %d\n", n)
	} else {
		m["z"] = 30
		println("inserted")
	}
}
