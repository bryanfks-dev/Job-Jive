package models

import (
	"db"
)

type DepartmentHead struct {
	DepartmentId int `json:"department_id"`
	ManagerId *int `json:"manager_id"`
}

func (departmentHead DepartmentHead) Insert() error {
	stmt := "INSERT INTO `department_heads` (Department_ID, Manager_ID) VALUES(?, ?)"

	_, err := db.Conn.Exec(stmt, 
			departmentHead.DepartmentId, departmentHead.ManagerId)

	return err
}