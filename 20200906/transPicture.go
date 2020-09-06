package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func changeImageExtension(imageFilePath string) error {
	imageFile, err := os.Open(imageFilePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return err
	}

	imageFilePathExt := filepath.Ext(imageFilePath)                                             // 変換前の画像ファイルの拡張子
	changedImageFileExt := ".png"                                                               // 変換後の画像ファイルの拡張子
	changedImageFilePath := strings.Trim(imageFilePath, imageFilePathExt) + changedImageFileExt // 変換後の画像ファイルのパス

	changedImageFile, err := os.Create(changedImageFilePath)
	if err != nil {
		return err
	}
	defer changedImageFile.Close()

	// pngに変換
	png.Encode(changedImageFile, img)

	return nil
}

func main() {
	// 画像形式オプションを取得（デフォルト値: jpg）
	imageFormatPointer := flag.String("imageFormat", "jpg", "image format")

	// ディレクトリオプションを取得（デフォルト値: images）
	dirPointer := flag.String("dir", "images", " search root directory")

	flag.Parse()

	// jpeg -> jpg
	if *imageFormatPointer == "jpeg" {
		*imageFormatPointer = "jpg"
	}

	imageExt := "." + *imageFormatPointer // 画像ファイルの拡張子
	imageFilePaths := make([]string, 0)   // 画像ファイルのパス群

	// 画像ファイルのパスを取得し, imageFilePathsに格納
	err := filepath.Walk(*dirPointer,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == imageExt {
				imageFilePaths = append(imageFilePaths, path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	if len(imageFilePaths) != 0 {
		for _, imageFilePath := range imageFilePaths {
			err := changeImageExtension(imageFilePath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
