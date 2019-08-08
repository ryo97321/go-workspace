package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

type Words struct {
	ID   int
	Word string
}

var testWords []string

func getWordsFromDB() {
	fmt.Print("MySQL root password : ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	dataSourceName := fmt.Sprintf("root:%s@/gotypingdb", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM WORDS ORDER BY RAND() LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var words Words
		err := rows.Scan(&words.ID, &words.Word)
		if err != nil {
			log.Fatal(err)
		}
		testWords = append(testWords, words.Word)
	}
}

func main() {
	var userInput string
	var answer string
	var processTimeSecond float64

	getWordsFromDB()

	stdin := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("start")
	fmt.Println()

	fmt.Println("********************")
	for i := 0; i < 10; i++ {
		answer = testWords[i]
		fmt.Printf("%d : %s\n", (i + 1), answer)

		startTime := time.Now()
		for {
			fmt.Print("Input : ")
			stdin.Scan()
			userInput = stdin.Text()

			if userInput == answer {
				fmt.Println()
				break
			}
		}
		endTime := time.Now()
		processTime := endTime.Sub(startTime)
		processTimeSecond += processTime.Seconds()

		if i != 9 {
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("********************")

	fmt.Printf("\nResult : %f秒\n", processTimeSecond)
	fmt.Println("end")
}
