package models

import (
	"db"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int
	FullName     string
	Email        string
	Password     string
	DateOfBirth  string
	Address      string
	NIK          string
	Photo        string
	Gender       string
	PhoneNumber  string
	DepartmentId *int
	FirstLogin   *string
}

func (user User) Get() ([]User, error) {
	stmt := "SELECT * FROM `users` ORDER BY User_ID DESC"

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

func (user User) GetUsingDepartmentId(department_id int) ([]User, error) {
	stmt := "SELECT * FROM `users` WHERE Department_ID = ?"

	row, err := db.Conn.Query(stmt, department_id)

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
	stmt := "SELECT * FROM `users` WHERE Department_ID = ? AND User_ID <> ? ORDER BY User_ID DESC"

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

func (user User) GetUsingEmail(email string) (User, error) {
	stmt := "SELECT * FROM `users` WHERE Email = ?"

	// Query result from user table with given email should
	// be returning 1 row, since the email value is unique
	err := db.Conn.QueryRow(stmt, email).
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

func (user User) Insert() (int, error) {
	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 11)

	if err != nil {
		return 0, err
	}

	user.Password = string(hash)

	stmt := "INSERT INTO `users` (Full_Name, Email, Password, Date_of_Birth, Address, NIK, Gender, Phone_Number, Department_ID, Photo) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	row, err := db.Conn.Exec(stmt,
		user.FullName,
		user.Email,
		user.Password,
		user.DateOfBirth,
		user.Address,
		user.NIK,
		user.Gender,
		user.PhoneNumber,
		user.DepartmentId,
		user.Photo)

	// Ensure no error inserting data
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()

	// Ensure getting last inserted id
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (user User) Update() error {
	stmt := "UPDATE `users` SET Full_Name = ?, Email = ?, Password = ?, Date_of_Birth = ?, Address = ?, NIK = ?, Gender = ?, Phone_Number = ?, Department_ID = ? WHERE User_ID = ?"

	_, err := db.Conn.Exec(stmt,
		user.FullName,
		user.Email,
		user.Password,
		user.DateOfBirth,
		user.Address,
		user.NIK,
		user.Gender,
		user.PhoneNumber,
		user.DepartmentId,
		user.Id)

	return err
}

func (user User) UpdateFistLogin(date string) error {
	stmt := "UPDATE `users` SET First_Login = ? WHERE User_Id = ?"

	_, err := db.Conn.Exec(stmt,
		date,
		user.Id)

	return err
}

func (user User) Delete() error {
	stmt := "DELETE FROM `users` WHERE User_ID = ?"

	_, err := db.Conn.Exec(stmt, user.Id)

	return err
}

func (user User) Search(query string) ([]User, error) {
	stmt := "SELECT u.* FROM `users` u LEFT JOIN `departments` d ON d.Department_ID = u.Department_ID WHERE CONCAT(u.Full_Name, '|', u.Email, '|', u.Phone_Number, '|', u.Date_of_Birth, '|', u.Gender, '|', COALESCE(d.Department_Name, '')) REGEXP ?"

	row, err := db.Conn.Query(stmt, query)

	// Ensure no error get user data
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

func (user User) SearchDepartmentEmployee(query string, department_id int, manager_id int) ([]User, error) {
	stmt := "SELECT * FROM `users` WHERE CONCAT(Full_Name, '|', Email, '|', Phone_Number, '|', Gender) REGEXP ? AND Department_ID = ? AND User_ID <> ?"

	row, err := db.Conn.Query(stmt, query, department_id, manager_id)

	// Ensure no error get user data
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
