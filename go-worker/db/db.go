package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *sql.DB
	err  error
)

func Connect(user string, password string, host string, port string, db_name string) error {
	// Open connection to database
	Conn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, db_name))

	return err
}

func ConnectionEstablished() bool {
	err := Conn.Ping()

	return err == nil
}
