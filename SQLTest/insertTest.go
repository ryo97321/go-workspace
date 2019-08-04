package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	stdio := bufio.NewScanner(os.Stdin)

	// input name
	var name string
	fmt.Print("Name : ")
	stdio.Scan()
	name = stdio.Text()

	// make query
	mySQLQuery := fmt.Sprintf("insert into users (Name) values (\"%s\");", name)

	// input password
	fmt.Print("MySQL root password : ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// make dataSourceName
	dataSourceName := fmt.Sprintf("root:%s@/gosample", password)

	// connect MySQL
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected.")

	// Exec insert Query
	result, err := db.Exec(mySQLQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted.")
	lastInsertID, _ := result.LastInsertId()
	fmt.Printf("Inserted ID : %d\n", lastInsertID)
}
