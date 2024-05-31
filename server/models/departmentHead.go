package models

import (
	"db"
)

type DepartmentHead struct {
	DepartmentId int  `json:"department_id"`
	ManagerId    *int `json:"manager_id"`
}

func (department_head DepartmentHead) Get() ([]DepartmentHead, error) {
	stmt := "SELECT * FROM `department_heads`"

	row, err := db.Conn.Query(stmt)

	// Ensure no error when fetching row
	if err != nil {
		return []DepartmentHead{}, err
	}

	defer row.Close()

	var department_heads []DepartmentHead

	for row.Next() {
		err := row.Scan(
			&department_head.DepartmentId, 
			&department_head.ManagerId)

		// Ensure no error when parsing row to struct
		if err != nil {
			return []DepartmentHead{}, err
		}

		department_heads = append(department_heads, department_head)
	}

	return department_heads, nil
}

func (department_head DepartmentHead) GetUsingDepartmentId(department_id int) (DepartmentHead, error) {
	stmt := "SELECT * FROM `department_heads` WHERE Department_ID = ?"

	row, err := db.Conn.Query(stmt, department_id)

	// Ensure no error when fetching row
	if err != nil {
		return DepartmentHead{}, err
	}

	defer row.Close()

	// Query result from department_head table with given department_id should
	// be returning 1 row, since the department_id value is unique
	if row.Next() {
		err := row.Scan(
			&department_head.DepartmentId, 
			&department_head.ManagerId)
		
		// Ensure no error when parsing row to struct
		if err != nil {
			return DepartmentHead{}, err
		}
	}

	return department_head, nil
}

func (department_head DepartmentHead) Insert() error {
	stmt := "INSERT INTO `department_heads` (Department_ID, Manager_ID) VALUES(?, ?)"

	_, err := db.Conn.Exec(stmt,
		department_head.DepartmentId, department_head.ManagerId)

	return err
}
