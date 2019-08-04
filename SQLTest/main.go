package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID   int
	Name string
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("MySQL root password : ")
	stdin.Scan()
	password := stdin.Text()

	dataSourceName := fmt.Sprintf("root:%s@/gosample", password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected.")

	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var users Users
		err := rows.Scan(&users.ID, &users.Name)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(users.ID, users.Name)
	}
}
