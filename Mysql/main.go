package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Names struct {
	Data string `json:"data"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Hello")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Print(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query("select * from names")

	if err != nil {
		panic(err.Error())

	}
	for results.Next() {
		var n Names

		err = results.Scan(&n.Data)
		if err != nil {
			panic(err.Error())

		}
		log.Printf(n.Data)
	}

}
