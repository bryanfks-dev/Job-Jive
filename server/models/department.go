package models

import (
	"db"
)

type Department struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
}

func (department Department) Get() ([]Department, error) {
	stmt := "SELECT * FROM `departments`"

	row, err := db.Conn.Query(stmt)

	defer row.Close()

	// Ensure no error when fetching records
	if err != nil {
		return []Department{}, err
	}

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

func (department Department) Insert() (error) {
	stmt := "INSERT INTO `departments` (Department_Name) VALUES(?)"

	_, err := db.Conn.Exec(stmt, department.Name)

	return err
}
