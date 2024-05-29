package models

import (
	"db"

	"golang.org/x/crypto/bcrypt"
)

type AdminCred struct {
	Username string
	Password string
}


type Admin struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (admin Admin) AddToDB() {
	if db.ConnectionEstablished() {
		// Hashing password
		hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 11)

		if err != nil {
			panic(err.Error())
		}

		admin.Password = string(hash);

		// Insert user into admin table
		stmt := "INSERT INTO `admins` VALUES('', ?, ?)"

		_, err = db.Conn.Exec(stmt, admin.Username, admin.Password)

		if err != nil {
			panic(err.Error())
		}
	}
}

func (admin Admin) GetHashedPassword(username string) string {
	stmt := "SELECT Password FROM `admins` WHERE Username= ?"

	row, err := db.Conn.Query(stmt, username)

	if err != nil {
		panic(err.Error())
	}

	defer row.Close()

	var user_pwd string

	// Query result from user table with given username should
	// be returning 1 row, since the username value is unique
	if row.Next() {
		err := row.Scan(&user_pwd)

		if err != nil {
			panic(err.Error())
		}
	}

	return user_pwd
}