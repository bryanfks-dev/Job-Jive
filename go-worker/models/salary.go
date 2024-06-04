package models

import "db"

type Salary struct {
	UserId        int     `json:"user_id"`
	IntialSalary  float64 `json:"initial_salary"`
	CurrentSalary float64 `json:"current_salary"`
}

func ResetCurrentSalary() error {
	stmt := "UPDATE `salaries` SET Current_Salary = Initial_Salary"

	_, err := db.Conn.Exec(stmt)

	return err
}