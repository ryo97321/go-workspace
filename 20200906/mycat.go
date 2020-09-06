package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var lineNumber int

func init() {
	lineNumber = 1
}

func writeFileToConsoleWithLineNumber(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	return nil
}

func writeFileToConsole(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
		lineNumber++
	}

	return nil
}

func main() {
	nOptionPointer := flag.Bool("n", false, "n flag")
	flag.Parse()

	filenames := flag.Args()

	if *nOptionPointer {
		for _, filename := range filenames {
			err := writeFileToConsoleWithLineNumber(filename)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if *nOptionPointer == false {
		for _, filename := range filenames {
			err := writeFileToConsole(filename)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
