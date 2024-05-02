package db

import (
	"fmt"
	"database/sql"

	"models"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func SearchUser(email string) models.User {
	stmt := "SELECT * FROM `users` WHERE email = ?"

	row, err := db.Query(stmt, email)

	defer row.Close()

	if err != nil {
		panic(err.Error())
	}

	// Query result from user table with given email should
	// be returning 1 row, since the email value is unique
	var user_data models.User

	if row.Next() {
		// Store row into user struct
		err := row.Scan(&user_data.Id, &user_data.FullName, &user_data.Email, 
			&user_data.Password, &user_data.ManagerId, &user_data.Address, 
			&user_data.NIK, &user_data.Gender, &user_data.PhoneNumber, 
			&user_data.DepartmentId, &user_data.FirstLogin)

		if err != nil {
			panic(err.Error())
		}
	}

	return user_data
}

func Connect(user string, password string, url string, port string, db_name string) {
	// Open connection to database
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, url, port, db_name))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Printf("Connected to database %s:%s (%s)\n", url, port, db_name)
}
