package models

import (
	// "cloud.google.com/go/civil"

	"db"
)

type User struct {
	Id           int    `json:"id"`
	FullName     string  `json:"full_name"`
	Photo        string  `json:"photo"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	DateOfBirth  *string `json:"date_of_birth"`
	Address      string  `json:"address"`
	NIK          string  `json:"nik"`
	Gender       string  `json:"gender"`
	PhoneNumber  string  `json:"phone_number"`
	DepartmentId int     `json:"department_id"`
	FirstLogin   *string `json:"first_login"`
}

func (user User) Insert() {
	stmt := "INSERT INTO `users` (User_ID, Full_Name, Email, Password, Manager_ID, Address, NIK, Gander, Phone_Number, Department_ID, First_Login) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Conn.Exec(stmt, 
		user.FullName, 
		user.Photo, 
		user.Email, 
		user.Password, 
		user.DateOfBirth, 
		user.Address, 
		user.NIK, 
		user.Gender, 
		user.PhoneNumber, 
		user.DepartmentId, 
		user.FirstLogin)

	if err != nil {
		panic(err.Error())
	}
}

func (user User) GetUsingEmail(email string) (User, error) {
	stmt := "SELECT * FROM `users` WHERE Email = ?"

	row, err := db.Conn.Query(stmt, email)

	if err != nil {
		return User{}, err
	}

	defer row.Close()

	// Query result from user table with given email should
	// be returning 1 row, since the email value is unique
	if row.Next() {
		err := row.Scan(
			&user.Id, 
			&user.FullName, 
			&user.Photo, 
			&user.Email, 
			&user.Password, 
			&user.DateOfBirth, 
			&user.Address, 
			&user.NIK, 
			&user.Gender, 
			&user.PhoneNumber, 
			&user.DepartmentId, 
			&user.FirstLogin)

		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func (user User) GetUsingId(id int) (User, error) {
	stmt := "SELECT * FROM `users` WHERE User_Id = ?"

	row, err := db.Conn.Query(stmt, id)

	if err != nil {
		return User{}, err
	}

	defer row.Close()

	// Query result from user table with given id should
	// be returning 1 row, since the id value is unique
	if row.Next() {
		err := row.Scan(
			&user.Id, 
			&user.FullName, 
			&user.Photo, 
			&user.Email, 
			&user.Password, 
			&user.DateOfBirth, 
			&user.Address, 
			&user.NIK, 
			&user.Gender, 
			&user.PhoneNumber, 
			&user.DepartmentId,  
			&user.FirstLogin)

		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}
