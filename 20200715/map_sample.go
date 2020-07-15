package main

import "fmt"

func main() {
	// var numbers map[string]int		 <- nilで初期化される
	numbers := make(map[string]int) // ゼロ値で初期化される

	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)

	delete(numbers, "one")
	fmt.Println(numbers)

	numbers2 := numbers
	numbers2["two"] = 22

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)
}
