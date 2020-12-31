package main

// func f(xp *int) {
// 	*xp = 100
// }

func main() {
	// var x int
	// f(&x)
	// fmt.Println(x)

	ns := []int{10, 20, 30}
	ns2 := ns
	ns[1] = 200
	println(ns[0], ns[1], ns[2])
	println(ns2[0], ns2[1], ns2[2])
}
