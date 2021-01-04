// Package transImg : 画像の拡張子を変換するパッケージ
package transImg

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"image/jpeg"
	"image/png"
)

// ImgObj : 画像の構造体
type ImgObj struct {
	InPath string // 変換前の画像のパス
	InExt  string // 変換前の画像の拡張子
	OutExt string // 変換後の画像の拡張子
}

func changeExt(path, inExt, outExt string) string {
	outputDir := filepath.Dir(path)
	outputFilenameWithoutExt := strings.Replace(filepath.Base(path), inExt, "", -1)
	outputExt := outExt
	outputFilename := outputFilenameWithoutExt + outputExt
	outputPath := filepath.Join(outputDir, outputFilename)

	return outputPath
}

// PngToJpg : pathの画像ファイルの拡張子を.pngから.jpgに変換する
func PngToJpg(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	var imgObj ImgObj
	imgObj.InPath = path
	imgObj.InExt = ".png"
	imgObj.OutExt = ".jpg"

	outputPath := changeExt(imgObj.InPath, imgObj.InExt, imgObj.OutExt)

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		return err
	}

	fmt.Println(path + " -> " + outputPath)
	return nil
}

// JpgToPng : pathの画像ファイルの拡張子を.jpgから.pngに変換する
func JpgToPng(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	var imgObj ImgObj
	imgObj.InPath = path
	imgObj.InExt = ".jpg"
	imgObj.OutExt = ".png"

	outputPath := changeExt(imgObj.InPath, imgObj.InExt, imgObj.OutExt)

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	fmt.Println(path + " -> " + outputPath)
	return nil
}
