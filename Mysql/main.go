package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Define the MySQL database connection parameters.
	dsn := "root:root@tcp(localhost:3306)/Hello"

	// Open a database connection.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database connection.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Query the database to select data from a table.
	query := "SELECT * FROM names"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the rows and print the data.
	for rows.Next() {
		var (
			column1Type string
			column2Type int
			// Define variables for each column in the table
		)
		if err := rows.Scan(&column1Type, &column2Type /* Assign variables for each column */); err != nil {
			log.Fatal(err)
		}

		// Process or print the data.
		fmt.Printf("Column1: %s, Column2: %d\n", column1Type, column2Type)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
