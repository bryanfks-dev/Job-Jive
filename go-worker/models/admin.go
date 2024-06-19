package models

import (
	"db"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id       int
	Username string
	Password string
}

func (admin Admin) GetUsingUsername(username string) (Admin, error) {
	stmt := "SELECT * FROM `admins` WHERE Username = ?"

	// Query result from admin table with given id should
	// be returning 1 row, since the username value is unique
	err := db.Conn.QueryRow(stmt, username).
		Scan(&admin.Id,
			&admin.Username,
			&admin.Password)

	return admin, err
}

func (admin Admin) GetUsingId(id int) (Admin, error) {
	stmt := "SELECT * FROM `admins` WHERE Admin_ID = ?"

	// Query result from admin table with given id should
	// be returning 1 row, since the username value is unique
	err := db.Conn.QueryRow(stmt, id).
		Scan(&admin.Id,
			&admin.Username,
			&admin.Password)

	return admin, err
}

func (admin Admin) Insert() error {
	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 11)

	if err != nil {
		return err
	}

	admin.Password = string(hash)

	// Insert user into admin table
	stmt := "INSERT INTO `admins` (Username, Password) VALUES(?, ?)"

	_, err = db.Conn.Exec(stmt, admin.Username, admin.Password)

	return err
}
