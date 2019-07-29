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

	var name string

	fmt.Print("Name : ")
	stdio.Scan()
	name = stdio.Text()

	mySQLQuery := fmt.Sprintf("insert into users (Name) values (\"%s\");", name)

	fmt.Print("MySQL root password : ")
	// stdio.Scan()
	// password := stdio.Text()

	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	dataSourceName := fmt.Sprintf("root:%s@/gosample", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected.")

	result, err := db.Exec(mySQLQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted.")
	lastInsertID, _ := result.LastInsertId()
	fmt.Printf("Inserted ID : %d\n", lastInsertID)

}
