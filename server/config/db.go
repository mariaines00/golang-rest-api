package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://asimov:password@localhost/factory?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
