package responses

import "models"

type DepartmentResponse struct {
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Manager models.User `json:"manager"`
}

func (department_response *DepartmentResponse) Create(department models.Department) error {
	department_head, err := 
		models.DepartmentHead{}.GetUsingDepartmentId(department.Id)
	
	// Ensure no error get department_head
	if err != nil {
		return err
	}

	department_response.Id = department.Id
	department_response.Name = department.Name

	if department_head.ManagerId != nil {
		user, err := 
			models.User{}.GetUsingId(*department_head.ManagerId)

		// Ensure no error fetching manager data
		if err != nil {
			return err
		}

		department_response.Manager = models.User{
			Id: user.Id,
			FullName: user.FullName,
			Email: user.Email,
			DateOfBirth: user.DateOfBirth,
			Address: user.Address,
			NIK: user.NIK,
			Gender: user.Gender,
			Photo: user.Photo,
			PhoneNumber: user.PhoneNumber,
			FirstLogin: user.FirstLogin,
		}
	}

	return nil
}
