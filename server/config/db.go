package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sql.DB

func init() {
	var err error
	dbURL := os.ExpandEnv("postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@database/$POSTGRES_DB?sslmode=disable")

	DB, err = sql.Open("postgres", dbURL)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
