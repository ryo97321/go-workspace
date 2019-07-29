package main

import (
	"database/sql"
	"fmt"
	"log"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

type Users struct {
	ID   int
	Name string
}

func main() {

	// input MySQL password
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

	// Exec select Query
	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	// view Result
	for rows.Next() {
		var users Users
		err := rows.Scan(&users.ID, &users.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(users.ID, users.Name)
	}
}
