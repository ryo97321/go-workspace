package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 4桁の数字にダブりがないか確認する関数
func isAnswerNumberDuplicate(answer int) bool {
	answerString := strconv.Itoa(answer)
	var checkList []int
	for i := 0; i < 10; i++ {
		checkList = append(checkList, 0)
	}

	for _, answerNumberString := range answerString {
		answerNumber, _ := strconv.Atoi(string(answerNumberString))
		checkList[answerNumber]++
	}

	for i := 0; i < 10; i++ {
		if checkList[i] > 1 {
			return true
		}
	}

	return false
}

// ユーザの入力と正解を比較しEAT数とBITE数を返す関数
func check(input, answer string) (nEat, nBite int) {

	for i := 0; i < 4; i++ {
		checkNumber, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < 4; j++ {
			number, _ := strconv.Atoi(string(answer[j]))
			if checkNumber == number && i == j {
				nEat++
			}
		}
	}

	return nEat, nBite

}

func main() {
	fmt.Println("Numer0n start.")

	rand.Seed(time.Now().UnixNano())

	var answer int
	for {
		answer = rand.Intn(9000) + 1000
		if isAnswerNumberDuplicate(answer) != true {
			break
		}
	}

	for {
		// ユーザの入力を受け取る
		stdin := bufio.NewScanner(os.Stdin)
		fmt.Print("Input Number : ")
		stdin.Scan()
		input := stdin.Text()

		// チェックする（関数を使用）
		nEat, nBite := check(input, strconv.Itoa(answer))

		// 結果を表示する（Eat, Bite）
		fmt.Printf("%dEAT-%dBITE\n\n", nEat, nBite)

		// 正解したら抜ける
		if input == strconv.Itoa(answer) {
			break
		}
	}

}
