package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var userInput string
	var answer string
	var testWords = []string{"apple", "banana", "orange", "red", "blue", "green", "black", "white", "dog", "cat"}

	stdin := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("start\n")

	fmt.Println("********************")
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		answer = testWords[rand.Intn(10)]
		fmt.Printf("%d : %s\n", (i + 1), answer)

		for {
			fmt.Print("Input : ")
			stdin.Scan()
			userInput = stdin.Text()

			if userInput == answer {
				fmt.Println()
				break
			}
		}
	}
	endTime := time.Now()
	processTime := endTime.Sub(startTime)
	fmt.Println("********************")

	fmt.Printf("\nResult : %f秒\n", processTime.Seconds())
	fmt.Println("end")
}
