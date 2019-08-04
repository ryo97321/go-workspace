package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Print("MySQL root password : ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	dataSourceName := fmt.Sprintf("root:%s@/gotypingdb", password)

	// connect DB
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected.")

	fp, err := os.Open("./wordsFolder/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// insert word -> MySQL
	var sqlQuery string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		word := scanner.Text()
		sqlQuery = fmt.Sprintf("insert into words (word) values (\"%s\");", word)
		_, err := db.Exec(sqlQuery)
		if err != nil {
			if err.(*mysql.MySQLError).Number == 1062 {
				continue
			} else {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Completed.")

}
