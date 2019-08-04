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
	var processTimeSecond float64

	stdin := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("start")
	fmt.Println()

	fmt.Println("********************")
	for i := 0; i < 10; i++ {
		answer = testWords[rand.Intn(10)]
		fmt.Printf("%d : %s\n", (i + 1), answer)

		startTime := time.Now()
		for {
			fmt.Print("Input : ")
			stdin.Scan()
			userInput = stdin.Text()

			if userInput == answer {
				fmt.Println()
				break
			}
		}
		endTime := time.Now()
		processTime := endTime.Sub(startTime)
		processTimeSecond += processTime.Seconds()

		if i != 9 {
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("********************")

	fmt.Printf("\nResult : %f秒\n", processTimeSecond)
	fmt.Println("end")
}
