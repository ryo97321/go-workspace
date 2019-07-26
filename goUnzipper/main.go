package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
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

		for _, f := range r.File {
			splitFileName := strings.Split(f.Name, "/")

			var mkdirPath string

			// mkdir
			if len(splitFileName) != 1 {
				for i := 0; i < len(splitFileName)-1; i++ {
					mkdirPath += splitFileName[i]
					if i != len(splitFileName)-2 {
						mkdirPath += "/"
					}
				}
				os.Mkdir(mkdirPath, f.Mode())
			}

			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer rc.Close()

			buf := make([]byte, f.UncompressedSize)
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				log.Fatal(err)
			}

			if err = ioutil.WriteFile(f.Name, buf, f.Mode()); err != nil {
				log.Fatal(err)
			}
		}
	}

}
