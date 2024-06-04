package models

import (
	"db"
)

type User struct {
	Id           int     `json:"id"`
	FullName     string  `json:"full_name"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	DateOfBirth  string  `json:"date_of_birth"`
	Address      string  `json:"address"`
	NIK          string  `json:"nik"`
	Photo        string  `json:"photo"`
	Gender       string  `json:"gender"`
	PhoneNumber  string  `json:"phone_number"`
	DepartmentId int     `json:"department_id"`
	FirstLogin   *string `json:"first_login"`
}

func (user User) GetUsingEmail(email string) (User, error) {
	stmt := "SELECT * FROM `users` WHERE Email = ?"

	// Query result from user table with given email should
	// be returning 1 row, since the email value is unique
	err := db.Conn.QueryRow(stmt, email).
		Scan(&user.Id,
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

	return user, err
}

func (user User) GetUsingId(id int) (User, error) {
	stmt := "SELECT * FROM `users` WHERE User_ID = ?"

	// Query result from user table with given id should
	// be returning 1 row, since the id value is unique
	err := db.Conn.QueryRow(stmt, id).
		Scan(&user.Id,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.DateOfBirth,
			&user.Address,
			&user.NIK,
			&user.Photo,
			&user.Gender,
			&user.PhoneNumber,
			&user.DepartmentId,
			&user.FirstLogin)

	return user, err
}

func (user User) GetUsers() ([]User, error) {
	stmt := "SELECT * FROM `users`"

	row, err := db.Conn.Query(stmt)

	if err != nil {
		return []User{}, err
	}

	defer row.Close()

	var users []User

	for row.Next() {
		err := row.Scan(&user.Id,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.DateOfBirth,
			&user.Address,
			&user.NIK,
			&user.Photo,
			&user.Gender,
			&user.PhoneNumber,
			&user.DepartmentId,
			&user.FirstLogin)

		if err != nil {
			return []User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (user User) GetEmployees(manager_id int, department_id int) ([]User, error) {
	stmt := "SELECT * FROM `users` WHERE Department_ID = ? AND User_ID <> ?"

	row, err := db.Conn.Query(stmt, department_id, manager_id)

	if err != nil {
		return []User{}, err
	}

	defer row.Close()

	var users []User

	for row.Next() {
		err := row.Scan(&user.Id,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.DateOfBirth,
			&user.Address,
			&user.NIK,
			&user.Photo,
			&user.Gender,
			&user.PhoneNumber,
			&user.DepartmentId,
			&user.FirstLogin)

		if err != nil {
			return []User{}, err
		}

		users = append(users, user)
	}

	return users, nil
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
