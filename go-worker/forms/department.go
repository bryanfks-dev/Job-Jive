package forms

import (
	"database/sql"
	"errors"
	"strings"

	"models"
)

type DepartmentForm struct {
	DepartmentName string `json:"department_name"`
	ManagerId      int    `json:"manager_id"`
}

var (
	ErrManagerIdExist = errors.New("user already become manager in other department")
)

func (department_form *DepartmentForm) Sanitize() {
	department_form.DepartmentName = strings.TrimSpace(department_form.DepartmentName)
}

func (department_form *DepartmentForm) Validate() (bool, error) {
	_, err :=
		models.DepartmentHead{}.GetUsingManagerId(department_form.ManagerId)

	// Ensure no error fetching department_head data
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	} else {
		return false, ErrManagerIdExist
	}

	return true, nil
}
