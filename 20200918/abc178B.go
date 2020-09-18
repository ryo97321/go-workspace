package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputs := scanner.Text()

	inputsSlice := strings.Split(inputs, " ")
	a, _ := strconv.Atoi(inputsSlice[0])
	b, _ := strconv.Atoi(inputsSlice[1])
	c, _ := strconv.Atoi(inputsSlice[2])
	d, _ := strconv.Atoi(inputsSlice[3])

	// a*c, a*d, b*c, b*d の結果を格納
	patterns := make([]int, 4)

	patterns[0] = a * c
	patterns[1] = a * d
	patterns[2] = b * c
	patterns[3] = b * d

	max := 0
	for i := 0; i < len(patterns); i++ {
		if i == 0 {
			max = patterns[i]
		}

		if patterns[i] > max {
			max = patterns[i]
		}
	}

	fmt.Println(max)
}
