package main

import "fmt"

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
func commandToBinaryString(command string) string {
	binaryString := ""

	for _, data := range command {
		binaryString += compensateEightBits(data)
	}

	return binaryString
}

func main() {
	command := "ls -la"
	fmt.Println(commandToBinaryString(command))
}
