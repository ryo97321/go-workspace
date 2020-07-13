package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// ファイルのオープン
	f, err := os.Open("file.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// ファイル読み込み用バッファ
	buffer := make([]byte, 10)

	// ファイル読み込み
	f.Read(buffer)

	// 改行を削除
	buffer = bytes.ReplaceAll(buffer, []byte{13, 10, 0, 0}, []byte(""))

	// stringに変換して表示
	s := string(buffer)
	fmt.Printf("s: %s\n", s)
}
