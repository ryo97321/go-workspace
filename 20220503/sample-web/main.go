package main

import (
	"database/sql"
	"fmt"

	_ "https://github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=4G3EzkW8a6 dbname=test sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username,dapartname,created) VALUES($1, $2, $3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("mori", "研究開発部門", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
