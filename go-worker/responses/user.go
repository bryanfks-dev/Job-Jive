package responses

import (
	"log"
	"models"
)

type UserResponse struct {
	Id          int               `json:"id"`
	FullName    string            `json:"full_name"`
	As          string            `json:"as"`
	Email       string            `json:"email"`
	Address     string            `json:"address"`
	BirthDate   string            `json:"birth_date"`
	PhoneNumber string            `json:"phone_number"`
	Gender      string            `json:"gender"`
	NIK         string            `json:"nik"`
	Department  models.Department `json:"department"`
	Photo       string            `json:"photo"`
	Salary      models.Salary     `json:"salary"`
	First_Login *string           `json:"first_login"`
}

func (user_response *UserResponse) Create(user models.User) error {
	department, err :=
		models.Department{}.GetUsingId(*user.DepartmentId)

	// Ensure no error fetching department
	if err != nil {
		log.Panic("Error get department: ", err.Error())

		return err
	}

	salary, err :=
		models.Salary{}.GetUsingUserId(user.Id)

	// Ensure no error fetching salary
	if err != nil {
		log.Panic("Error get salary: ", err.Error())

		return err
	}

	department_head, err :=
		models.DepartmentHead{}.GetUsingDepartmentId(*user.DepartmentId)

	// Ensure no error get department_head
	if err != nil {
		log.Panic("Error get department_head: ", err.Error())

		return err
	}

	// Decide whether user as manage or employee
	as := "Employee"

	if department_head.ManagerId != nil && user.Id == *department_head.ManagerId {
		as = "Manager"
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
	user_response.Department = department
	user_response.Photo = user.Photo
	user_response.Salary = models.Salary{
		Initial: salary.Initial,
		Current: salary.Current,
	}
	user_response.First_Login = user.FirstLogin

	return nil
}
