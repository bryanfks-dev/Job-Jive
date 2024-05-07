package db

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn  *sql.DB
	err error
)

func Connect(user string, password string, url string, port string, db_name string) {
	// Open connection to database
	Conn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, url, port, db_name))

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Connected to database %s:%s (%s)\n", url, port, db_name)
}

func ConnectionEstablished() bool {
	err := Conn.Ping()

	return err == nil
}