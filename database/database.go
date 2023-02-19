package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

var DB *sql.DB

func Connect() error {
	var err error
	// Todo use config file to get from dotenv
	p := "5432"
	host := "localhost"
	user := "ticket_adm"
	password := "password"
	dbname := "ticket"

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
