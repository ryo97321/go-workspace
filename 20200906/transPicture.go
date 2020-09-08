package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	myimage "github.com/ryo97321/image"
)

func main() {
	// ディレクトリオプションを取得（デフォルト値: images）
	dirPointer := flag.String("dir", "images", "search root directory")

	// 変換前のファイルの拡張子を取得
	srcFileExtPointer := flag.String("srcFileExt", "jpg", "src image file Extension")

	// 変換後のファイルの拡張子を取得
	dstFileExtPointer := flag.String("dstFileExt", "png", "dest image file Extension")

	flag.Parse()

	// jpeg -> jpg
	if *srcFileExtPointer == "jpeg" {
		*srcFileExtPointer = "jpg"
	}
	if *dstFileExtPointer == "jpeg" {
		*dstFileExtPointer = "jpg"
	}

	srcFileExt := "." + *srcFileExtPointer // 変換前の画像ファイルの拡張子
	dstFileExt := "." + *dstFileExtPointer // 変換後の画像ファイルの拡張子
	imageFilePaths := make([]string, 0)    // 画像ファイルのパス群

	// 画像ファイルのパスを取得し, imageFilePathsに格納
	err := filepath.Walk(*dirPointer,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == srcFileExt {
				imageFilePaths = append(imageFilePaths, path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	// 画像の拡張子を変換
	if len(imageFilePaths) != 0 {
		for _, imageFilePath := range imageFilePaths {
			err := myimage.ChangeImageExtension(imageFilePath, srcFileExt, dstFileExt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
