package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	var pre_z float64
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v : %v\n", i, z)

		if i == 0 {
			pre_z = z
			continue
		}

		if pre_z >= z {
			if pre_z-z < 0.0000000001 {
				break
			}
		} else if pre_z < z {
			if z-pre_z < 0.0000000001 {
				break
			}
		}

		pre_z = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
