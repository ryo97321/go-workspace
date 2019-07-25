package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func searchZip(dirName string) []string {
	files, _ := ioutil.ReadDir(dirName)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, searchZip(filepath.Join(dirName, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dirName, file.Name()))
	}

	return paths
}

func main() {

	zipFileNames := searchZip("./")

	for _, name := range zipFileNames {
		fmt.Println(name)
	}

	// files, _ := ioutil.ReadDir("./")

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// if file.IsDir() != true {
	// 	content, _ := ioutil.ReadFile(file.Name())
	// 	fmt.Printf("Content : %s\n", string(content))
	// }
	// }
}
