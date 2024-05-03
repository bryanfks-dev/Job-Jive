package models

import (
	"time"

	// "cloud.google.com/go/civil"

	"db"
)

type DateType time.Time

func (t DateType) String() string {
	return time.Time(t).String()
}

type User struct {
	Id           int        `json:"id"`
	FullName     string     `json:"full_name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	ManagerId    *int       `json:"manager_id"`
	Address      string     `json:"address"`
	NIK          string     `json:"nik"`
	Gender       string     `json:"gender"`
	PhoneNumber  string     `json:"phone_number"`
	DepartmentId int        `json:"department_id"`
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

func (user User) Search(email string) User {
	stmt := "SELECT * FROM `users` WHERE email = ?"

	row, err := db.Conn.Query(stmt, email)

	if err != nil {
		panic(err.Error())
	}

	defer row.Close()

	// Query result from user table with given email should
	// be returning 1 row, since the email value is unique
	var user_data User

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
