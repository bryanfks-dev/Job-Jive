package models

import (
	"db"
)

type Department struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (department Department) Get() ([]Department, error) {
	stmt := "SELECT * FROM `departments`"

	row, err := db.Conn.Query(stmt)

	// Ensure no error when fetching records
	if err != nil {
		return []Department{}, err
	}

	defer row.Close()

	var departments []Department

	// Iterate through records and insert to array
	for row.Next() {
		err := row.Scan(
			&department.Id,
			&department.Name)

		// Ensure no error when decode rows into structs
		if err != nil {
			return []Department{}, err
		}

		departments = append(departments, department)
	}

	return departments, nil
}

func (department Department) GetUsingId(id int) (Department, error) {
	stmt := "SELECT * FROM `departments` WHERE Department_ID = ?"

	// Query result from department table with given id should
	// be returning 1 row, since the id value is unique
	err := db.Conn.QueryRow(stmt, id).
		Scan(&department.Id,
			&department.Name)

	return department, err
}

func (department Department) Insert() (int, error) {
	stmt := "INSERT INTO `departments` (Department_Name) VALUES(?)"

	row, err := db.Conn.Exec(stmt, department.Name)

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

func (department Department) Update() error {
	stmt := "UPDATE `departments` SET Department_Name = ? WHERE Department_ID = ?"

	_, err := db.Conn.Exec(stmt, department.Name, department.Id)

	return err
}

func (department Department) Delete() error {
	stmt := "DELETE FROM `departments` WHERE Department_ID = ?"

	_, err := db.Conn.Exec(stmt, department.Id)

	return err
}
