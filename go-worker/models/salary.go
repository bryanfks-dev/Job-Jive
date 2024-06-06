package models

import "db"

type Salary struct {
	UserId        int     `json:"user_id"`
	InitialSalary  float64 `json:"initial_salary"`
	CurrentSalary float64 `json:"current_salary"`
}

func (salary Salary) Insert() error {
	stmt := "INSERT INTO `salaries` (User_ID, Initial_Salary, Current_Salary) VALUES(?, ?, ?)"

	_, err := 
		db.Conn.Exec(stmt, salary.UserId, salary.InitialSalary, salary.CurrentSalary)

	return err
}

func ResetCurrentSalary() error {
	stmt := "UPDATE `salaries` SET Current_Salary = Initial_Salary"

	_, err := db.Conn.Exec(stmt)

	return err
}