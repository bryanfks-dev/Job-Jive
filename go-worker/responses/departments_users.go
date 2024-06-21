package responses

import (
	"log"
	"models"
)

type DepartmentsUsersResponse struct {
	Department DepartmentResponse `json:"department"`
	Users      []UserResponse     `json:"users"`
}

func (departments_users_response *DepartmentsUsersResponse) Create(department models.Department) error {
	users, err :=
		models.User{}.GetUsingDepartmentId(department.Id)

	if err != nil {
		log.Panic("Error get departments users: ", err.Error())

		return err
	}

	var department_response DepartmentResponse

	err = department_response.Create(department)

	if err != nil {
		return err
	}

	var user_response UserResponse

	var users_response []UserResponse

	for _, user := range users {
		err = user_response.Create(user)

		if err != nil {
			return err
		}

		users_response = append(users_response, user_response)
	}

	departments_users_response.Department = department_response
	departments_users_response.Users = users_response

	return nil
}
