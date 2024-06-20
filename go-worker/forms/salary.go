package forms

import "errors"

type SalaryForm struct {
	Initial float64 `json:"initial_salary"`
	Current float64 `json:"current_salary"`
}

var (
	ErrInvalidInitialSalary = errors.New("invalid initial salary value")
	ErrInvalidCurrentSalary = errors.New("invalid current salary value")
)

func (salary_form SalaryForm) Validate() (bool, error) {
	if salary_form.Initial <= 0 {
		return false, ErrInvalidInitialSalary
	}

	if salary_form.Current <= 0 {
		return false, ErrInvalidCurrentSalary
	}

	return true, nil
}
