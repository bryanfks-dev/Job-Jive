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

	row, err := db.Conn.Query(stmt, id)

	// Ensure no error when fetching records
	if err != nil {
		return Department{}, err
	}

	defer row.Close()

	// Query result from department table with given id should
	// be returning 1 row, since the id value is unique
	if row.Next() {
		err := row.Scan(
			&department.Id, 
			&department.Name)

		// Ensure no error when parsing row
		if err != nil {
			return Department{}, err
		}
	}

	return department, nil
}

func (department Department) Insert() error {
	stmt := "INSERT INTO `departments` (Department_Name) VALUES(?)"

	_, err := db.Conn.Exec(stmt, department.Name)

	return err
}
