package models

import (
	// "cloud.google.com/go/civil"

	"db"
)

type UserCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type User struct {
	Id           uint    `json:"id"`
	FullName     string  `json:"full_name"`
	PhotoPath    string  `json:"photo_path"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	ManagerId    *int    `json:"manager_id"`
	Address      string  `json:"address"`
	NIK          string  `json:"nik"`
	Gender       string  `json:"gender"`
	PhoneNumber  string  `json:"phone_number"`
	DepartmentId int     `json:"department_id"`
	FirstLogin   *string `json:"first_login"`
}

func (user User) Insert() {
	stmt := "INSERT INTO `users` (User_ID, Full_Name, Email, Password, Manager_ID, Address, NIK, Gander, Phone_Number, Department_ID, First_Login) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Conn.Exec(stmt, user.FullName, user.Email, user.Password, user.ManagerId,
		user.Address, user.NIK, user.Gender, user.PhoneNumber, user.DepartmentId, user.FirstLogin)

	if err != nil {
		panic(err.Error())
	}
}

func (user User) GetHashedPassword(email string) string {
	stmt := "SELECT Password FROM `users` WHERE Email = ?"

	row, err := db.Conn.Query(stmt, email)

	if err != nil {
		panic(err.Error())
	}

	defer row.Close()

	var user_pwd string

	// Query result from user table with given email should
	// be returning 1 row, since the email value is unique
	if row.Next() {
		err := row.Scan(&user_pwd)

		if err != nil {
			panic(err.Error())
		}
	}

	return user_pwd
}
