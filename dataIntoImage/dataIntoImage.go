package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"os"
	"strconv"
)

// 先頭に0を追加し, 8ビットに調整する
func compensateEightBits(bits int32) string {
	bitsString := fmt.Sprintf("%b", bits)
	lenBits := len(bitsString)

	compensatedBits := fmt.Sprintf("%b", bits)
	numOfZeroToAdd := 8 - lenBits
	for i := 0; i < numOfZeroToAdd; i++ {
		compensatedBits = "0" + compensatedBits
	}

	return compensatedBits
}

// コマンドをバイナリの文字列にする
func toBinaryString(command string) string {
	binaryString := ""

	for _, data := range command {
		binaryString += compensateEightBits(data)
	}

	return binaryString
}

func main() {
	command := "ls -la"
	binariedCommand := toBinaryString(command)
	binariedCommandLength := len(binariedCommand)

	// 元画像を作成
	originalImage_file, err := os.Open("cat.jpg")
	defer originalImage_file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 読み込み用に変換
	originalImage_image, _, _ := image.Decode(originalImage_file)
	originalImage_Bounds := originalImage_image.Bounds()
	originalImage_RGBA := image.NewRGBA(originalImage_Bounds)
	originalImage_Rect := originalImage_RGBA.Rect

	var imageHeght = originalImage_Rect.Max.Y // 画像の高さ
	var imageWidth = originalImage_Rect.Max.X // 画像の幅

	// データ書き込み用の画像の宣言
	exportImage_RGBA := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeght))
	exportImage_Rect := exportImage_RGBA.Rect

	// データ読み込み & 書き込み
	var dataWriteCount = 0
	for h := exportImage_Rect.Min.Y; h < exportImage_Rect.Max.Y; h++ {
		for v := exportImage_Rect.Min.X; v < exportImage_Rect.Max.X; v++ {
			r, g, b, _ := originalImage_image.At(v, h).RGBA()
			r, g, b = r>>8, g>>8, b>>8
			r_str := compensateEightBits(int32(r))
			if dataWriteCount < binariedCommandLength {
				lastBitChangedR_str := r_str[:7] + string(binariedCommand[dataWriteCount])
				lastBitChangedR_int64, _ := strconv.ParseInt(lastBitChangedR_str, 2, 0)
				lastBitChangedR_uint8 := uint8(lastBitChangedR_int64)
				exportImage_RGBA.Set(v, h, color.RGBA{lastBitChangedR_uint8, uint8(g), uint8(b), 0})
				dataWriteCount++
			} else {
				exportImage_RGBA.Set(v, h, color.RGBA{uint8(r), uint8(g), uint8(b), 0})
			}
		}
	}

	// 出力用画像ファイルを作成
	exportFile, err := os.Create("cat2.jpg")
	defer exportFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 出力用画像ファイルにデータをエンコード
	if err := jpeg.Encode(exportFile, exportImage_RGBA, &jpeg.Options{100}); err != nil {
		log.Fatal(err)
	}
}
