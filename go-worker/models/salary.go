package models

import "db"

type Salary struct {
	UserId  int
	Initial float64
	Current float64
}

func (salary Salary) GetUsingUserId(user_id int) (Salary, error) {
	stmt := "SELECT * FROM `salaries` WHERE User_ID = ?"

	err := db.Conn.QueryRow(stmt, user_id).
		Scan(&salary.UserId,
			&salary.Initial,
			&salary.Current)

	return salary, err
}

func (salary Salary) Insert() error {
	stmt := "INSERT INTO `salaries` (User_ID, Initial_Salary, Current_Salary) VALUES(?, ?, ?)"

	_, err :=
		db.Conn.Exec(stmt, salary.UserId, salary.Initial, salary.Current)

	return err
}

func (salary Salary) Update() error {
	stmt := "UPDATE `salaries` SET Initial_Salary = ?, Current_Salary = ? WHERE User_ID = ?"

	_, err := db.Conn.Exec(stmt, salary.Initial, salary.Current, salary.UserId)

	return err
}

func ResetCurrentSalary() error {
	stmt := "UPDATE `salaries` SET Current_Salary = Initial_Salary"

	_, err := db.Conn.Exec(stmt)

	return err
}
