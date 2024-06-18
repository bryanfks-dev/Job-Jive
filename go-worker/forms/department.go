package forms

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"models"
)

type DepartmentForm struct {
	DepartmentName string `json:"department_name"`
	ManagerId      int    `json:"manager_id"`
}

var (
	ErrDepartmentNameExist = errors.New("department name already exist")
	ErrManagerIdExist = errors.New("user already become a manager in other department")
)

func (department_form *DepartmentForm) Sanitize() {
	department_form.DepartmentName = strings.TrimSpace(department_form.DepartmentName)
}

func (department_form *DepartmentForm) ValidateCreate() (bool, error) {
	_, err :=
		models.Department{}.GetUsingDepartmentName(department_form.DepartmentName)

	// Ensure no error fetching department
	if err != nil {
		if err != sql.ErrNoRows {
			log.Panic("Error get department using name: ", err.Error())

			return false, err
		}
	} else {
		return false, ErrDepartmentNameExist
	}

	return true, nil
}

func (department_form *DepartmentForm) ValidateUpdate() (bool, error) {
	_, err :=
		models.DepartmentHead{}.GetUsingManagerId(department_form.ManagerId)

	// Ensure no error fetching department_head data
	if err != nil {
		if err != sql.ErrNoRows {
			log.Panic("Error get department_head: ", err.Error())

			return false, err
		}
	} else {
		return false, ErrManagerIdExist
	}

	return true, nil
}
