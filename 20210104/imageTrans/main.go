package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"projects/go-workspace/20210104/imageTrans/transImg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Root directory : ")
	scanner.Scan()
	dirPath := scanner.Text()

	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".jpg" {
				transImg.JpgToPng(path)
			} else if filepath.Ext(path) == ".png" {
				transImg.PngToJpg(path)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
}
