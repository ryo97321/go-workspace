package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "root@/gosample")
	log.Println("Connected to mysql.")

	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.Name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(person.ID, person.Name)
	}
}
