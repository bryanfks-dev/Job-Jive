package responses

import (
	"database/sql"
	"log"
	"models"
)

type UserResponse struct {
	Id          int                 `json:"id"`
	FullName    string              `json:"full_name,omitempty"`
	As          string              `json:"as,omitempty"`
	Email       string              `json:"email,omitempty"`
	Address     string              `json:"address,omitempty"`
	BirthDate   string              `json:"birth_date,omitempty"`
	PhoneNumber string              `json:"phone_number,omitempty"`
	Gender      string              `json:"gender,omitempty"`
	NIK         string              `json:"nik,omitempty"`
	Department  *DepartmentResponse `json:"department,omitempty"`
	Photo       string              `json:"photo,omitempty"`
	Salary      SalaryResponse      `json:"salary,omitempty"`
	FirstLogin  *string             `json:"first_login,omitempty"`
}

func (user_response *UserResponse) Create(user models.User) error {
	salary, err :=
		models.Salary{}.GetUsingUserId(user.Id)

	// Ensure no error fetching salary
	if err != nil {
		log.Panic("Error get salary: ", err.Error())

		return err
	}

	// Decide whether user as manage or employee
	as := "Employee"

	if user.DepartmentId != nil {
		department, err :=
			models.Department{}.GetUsingId(*user.DepartmentId)

		// Ensure no error fetching department
		if err != nil && err != sql.ErrNoRows {
			log.Panic("Error get department: ", err.Error())

			return err
		}

		department_head, err :=
			models.DepartmentHead{}.GetUsingDepartmentId(*user.DepartmentId)

		// Ensure no error get department_head
		if err != nil {
			log.Panic("Error get department_head: ", err.Error())

			return err
		}

		if department_head.ManagerId != nil && user.Id == *department_head.ManagerId {
			as = "Manager"
		}

		user_response.Department = &DepartmentResponse{
			Id:   department.Id,
			Name: department.Name,
		}
	}

	user_response.Id = user.Id
	user_response.FullName = user.FullName
	user_response.As = as
	user_response.Email = user.Email
	user_response.Address = user.Address
	user_response.BirthDate = user.DateOfBirth
	user_response.PhoneNumber = user.PhoneNumber
	user_response.Gender = user.Gender
	user_response.NIK = user.NIK
	user_response.Photo = user.Photo
	user_response.Salary = SalaryResponse{
		Initial: salary.Initial,
		Current: salary.Current,
	}
	user_response.FirstLogin = user.FirstLogin

	return nil
}
