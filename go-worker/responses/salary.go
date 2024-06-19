package responses

import (
	"log"
	"models"
)

type SalaryResponse struct {
	UserId  int     `json:"user_id"`
	Initial float64 `json:"initial_salary"`
	Current float64 `json:"current_salary"`
}

func (salary_response *SalaryResponse) Create(salary models.Salary) error {
	salary, err :=
		models.Salary{}.GetUsingUserId(salary.UserId)

	// Ensure no error fetching salary
	if err != nil {
		log.Panic("Error get salary: ", err.Error())

		return err
	}

	salary_response.UserId = salary.UserId
	salary_response.Initial = salary.Initial
	salary_response.Current = salary.Current

	return nil
}
