package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	destDir := "./output"
	zipFileName := "./test.zip"

	r, err := zip.OpenReader(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		if f.FileInfo().IsDir() {
			path := filepath.Join(destDir, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize)
			_, err := io.ReadFull(rc, buf)
			if err != nil {
				log.Fatal(err)
			}

			path := filepath.Join(destDir, f.Name)
			if err = ioutil.WriteFile(path, buf, f.Mode()); err != nil {
				log.Fatal(err)
			}
		}
	}
}
