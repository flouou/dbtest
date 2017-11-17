package main

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "<User>:<Password>@/<DB>")
	if err != nil {
		panic(err.Error)
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO user VALUES( ?, ?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT email FROM user WHERE username = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec("username"+strconv.Itoa(i), "mail@mail.test"+strconv.Itoa(i), "password"+strconv.Itoa(i)) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}
