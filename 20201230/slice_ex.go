package main

func main() {
	n := []int{19, 86, 1, 12}

	var sum int = 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}

	println(sum)
}
