package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func searchFiles(dirName string) []string {
	files, _ := ioutil.ReadDir(dirName)
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, searchFiles(filepath.Join(dirName, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dirName, file.Name()))
	}
	return paths
}

func searchZips(files []string) []string {
	var zips []string
	for _, file := range files {
		if strings.HasSuffix(file, ".zip") {
			zips = append(zips, file)
		}
	}
	return zips
}

func main() {
	files := searchFiles("./")
	zipFiles := searchZips(files)

	for _, zipFile := range zipFiles {
		r, err := zip.OpenReader(zipFile)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		fmt.Printf("*****%s*****\n", zipFile)
		for _, f := range r.File {
			fmt.Println(f.Name)
		}
		fmt.Println()
	}

}
