package models

import (
	"db"
)

type Department struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (department Department) Insert() {
	stmt := "INSERT INTO `departments` (Department_Name) VALUES(?)"

	_, err := db.Conn.Exec(stmt, department.Name)

	if err != nil {
		panic(err.Error())
	}
}