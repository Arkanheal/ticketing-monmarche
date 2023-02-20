package database

import (
	"database/sql"
	"fmt"
	"strconv"

    "tickets/config"
)

var DB *sql.DB

func Connect() error {
	var err error

	p := config.Config("DB_PORT")
	host := config.Config("DB_HOST")
	user := config.Config("POSTGRES_USER")
	password := config.Config("POSTGRES_PASSWORD")
	dbname := config.Config("POSTGRES_DB")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Mauvais port")
	}

	DB, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			password,
			dbname,
		),
	)

    if err != nil {
        return err
    }

    if err = DB.Ping(); err != nil {
        return err
    }

    CreateTables()
    fmt.Println("Connection Opened to Database")
	return nil
}
