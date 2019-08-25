package main

import (
	"bufio"
	"fmt"
	"image"
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

func main() {
	var command string
	var commandLength int
	var binariedCommand string
	var binariedCommandLength int

	// もとの文字列の長さを取得
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("Input command Length : ")
	stdin.Scan()
	input := stdin.Text()
	commandLength, _ = strconv.Atoi(input)
	binariedCommandLength = commandLength * 8

	// データが入っている画像を作成
	dataContainImage_file, err := os.Open("cat2.jpg")
	defer dataContainImage_file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 読み込み用に変換
	dataContainImage_image, _, _ := image.Decode(dataContainImage_file)
	dataContainImage_Bounds := dataContainImage_image.Bounds()
	dataContainImage_RGBA := image.NewRGBA(dataContainImage_Bounds)
	dataContainImage_Rect := dataContainImage_RGBA.Rect

	var isAllDataRead = false // すべてのデータを読み込んだか
	var dataReadCount = 0     // データを読み込んだ回数

	// データ読み込み
	for h := dataContainImage_Rect.Min.Y; h < dataContainImage_Rect.Max.Y; h++ {
		for v := dataContainImage_Rect.Min.X; v < dataContainImage_Rect.Max.X; v++ {
			dataReadCount++
			if dataReadCount <= binariedCommandLength {
				r, _, _, _ := dataContainImage_image.At(v, h).RGBA()
				r = r >> 8
				r_str := compensateEightBits(int32(r))
				binariedCommand += string(r_str[7])
			} else {
				isAllDataRead = true
				break
			}
		}
		if isAllDataRead {
			break
		}
	}

	// 変数に格納
	for i := 0; i < len(binariedCommand); i += 8 {
		binary_str := binariedCommand[i : i+8]
		binary_int64, _ := strconv.ParseInt(binary_str, 2, 0)
		command += string(binary_int64)
	}

	// 結果を表示
	fmt.Println(command)
}
