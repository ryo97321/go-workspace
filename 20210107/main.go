package main

import "errors"

func main() {
	err := errors.New("Error")
	println(err)
}
