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
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (admin Admin) Insert() {
	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 11)

	if err != nil {
		panic(err.Error())
	}

	admin.Password = string(hash)

	// Insert user into admin table
	stmt := "INSERT INTO `admins` (Username, Password) VALUES(?, ?)"

	_, err = db.Conn.Exec(stmt, admin.Username, admin.Password)

	if err != nil {
		panic(err.Error())
	}
}

func (admin Admin) GetUsingUsername(username string) (Admin, error) {
	stmt := "SELECT * FROM `admins` WHERE Username= ?"

	row, err := db.Conn.Query(stmt, username)

	if err != nil {
		return admin, err
	}

	defer row.Close()

	// Query result from user table with given username should
	// be returning 1 row, since the username value is unique
	if row.Next() {
		err := row.Scan(admin.Id, admin.Username, admin.Password)

		if err != nil {
			return admin, err
		}
	}

	return admin, nil
}
