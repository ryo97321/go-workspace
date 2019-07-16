package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	spinnerCharSetIndex := 31
	s := spinner.New(spinner.CharSets[spinnerCharSetIndex], 100*time.Millisecond)

	maxCharLength := 0
	for i := 0; i < len(spinner.CharSets[spinnerCharSetIndex]); i++ {
		if len(spinner.CharSets[spinnerCharSetIndex][i]) > maxCharLength {
			maxCharLength = len(spinner.CharSets[spinnerCharSetIndex][i])
		}
	}
	blanc := ""
	for i := 0; i < maxCharLength; i++ {
		blanc += " "
	}

	fmt.Println("start")

	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()

	fmt.Printf("\r%s\r", blanc) // clear
	fmt.Println("done")
}
