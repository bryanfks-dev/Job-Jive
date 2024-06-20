package responses

import (
	"log"
	"models"
)

type DepartmentResponse struct {
	Id        int             `json:"id"`
	Name      string          `json:"name,omitempty"`
	Manager   *UserResponse   `json:"manager"`
	IsManager *bool           `json:"is_manager,omitempty"`
	Employees *[]UserResponse `json:"employees,omitempty"`
}

func (department_response *DepartmentResponse) Create(department models.Department) error {
	department_head, err :=
		models.DepartmentHead{}.GetUsingDepartmentId(department.Id)

	// Ensure no error get department_head
	if err != nil {
		log.Panic("Error get department_head: ", err.Error())

		return err
	}

	department_response.Id = department.Id
	department_response.Name = department.Name

	if department_head.ManagerId != nil {
		manager, err :=
			models.User{}.GetUsingId(*department_head.ManagerId)

		// Ensure no error fetching manager data
		if err != nil {
			log.Panic("Error get manager user: ", err.Error())

			return err
		}

		manager_salary, err :=
			models.Salary{}.GetUsingUserId(manager.Id)

		// Ensure no error fetching manager salary
		if err != nil {
			log.Panic("Error get manager salary: ", err.Error())

			return err
		}

		department_response.Manager = &UserResponse{
			Id:        manager.Id,
			FullName:  manager.FullName,
			Email:     manager.Email,
			BirthDate: manager.DateOfBirth,
			Address:   manager.Address,
			NIK:       manager.NIK,
			Gender:    manager.Gender,
			Photo:     manager.Photo,
			Salary: SalaryResponse{
				Initial: manager_salary.Initial,
				Current: manager_salary.Current,
			},
			PhoneNumber: manager.PhoneNumber,
			FirstLogin:  manager.FirstLogin,
		}
	}

	return nil
}

func (department_response *DepartmentResponse) CreateUsers(department models.Department) error {
	department_head, err :=
		models.DepartmentHead{}.GetUsingDepartmentId(department.Id)

	// Ensure no error get department_head
	if err != nil {
		log.Panic("Error get department_head: ", err.Error())

		return err
	}

	department_response.Id = department.Id
	department_response.Name = department.Name

	if department_head.ManagerId != nil {
		// Fetch manager data
		manager, err :=
			models.User{}.GetUsingId(*department_head.ManagerId)

		// Ensure no error fetching manager data
		if err != nil {
			log.Panic("Error get manager user: ", err.Error())

			return err
		}

		manager_salary, err :=
			models.Salary{}.GetUsingUserId(manager.Id)

		// Ensure no error fetching manager salary
		if err != nil {
			log.Panic("Error get manager salary: ", err.Error())

			return err
		}

		department_response.Manager = &UserResponse{
			Id:        manager.Id,
			FullName:  manager.FullName,
			Email:     manager.Email,
			BirthDate: manager.DateOfBirth,
			Address:   manager.Address,
			NIK:       manager.NIK,
			Gender:    manager.Gender,
			Photo:     manager.Photo,
			Salary: SalaryResponse{
				Initial: manager_salary.Initial,
				Current: manager_salary.Current,
			},
			PhoneNumber: manager.PhoneNumber,
			FirstLogin:  manager.FirstLogin,
		}
	}

	manager_id := 0

	if department_head.ManagerId != nil {
		manager_id = *department_head.ManagerId
	}

	// Fetch employee data
	employees, err :=
		models.User{}.GetEmployees(manager_id, department.Id)

	if err != nil {
		log.Panic("Error get employees: ", err.Error())

		return err
	}

	var user_responses []UserResponse

	for _, employee := range employees {
		var user_response UserResponse

		salary, err :=
			models.Salary{}.GetUsingUserId(employee.Id)

		// Ensure no error get employee salary
		if err != nil {
			log.Panic("Error get employee salary: ", err.Error())

			return err
		}

		user_response.Id = employee.Id
		user_response.FullName = employee.FullName
		user_response.Email = employee.Email
		user_response.Address = employee.Address
		user_response.BirthDate = employee.DateOfBirth
		user_response.PhoneNumber = employee.PhoneNumber
		user_response.Gender = employee.Gender
		user_response.NIK = employee.NIK
		user_response.Photo = employee.Photo
		user_response.Salary = SalaryResponse{
			Initial: salary.Initial,
			Current: salary.Current,
		}
		user_response.FirstLogin = employee.FirstLogin

		user_responses = append(user_responses, user_response)
	}

	department_response.Employees = &user_responses

	return nil
}

func (department_response *DepartmentResponse) CreateUsersSearch(query string, department models.Department) error {
	department_head, err :=
		models.DepartmentHead{}.GetUsingDepartmentId(department.Id)

	// Ensure no error get department_head
	if err != nil {
		log.Panic("Error get department_head: ", err.Error())

		return err
	}

	department_response.Id = department.Id
	department_response.Name = department.Name

	if department_head.ManagerId != nil {
		// Fetch manager data
		manager, err :=
			models.User{}.GetUsingId(*department_head.ManagerId)

		// Ensure no error fetching manager data
		if err != nil {
			log.Panic("Error get manager user: ", err.Error())

			return err
		}

		manager_salary, err :=
			models.Salary{}.GetUsingUserId(manager.Id)

		// Ensure no error fetching manager salary
		if err != nil {
			log.Panic("Error get manager salary: ", err.Error())

			return err
		}

		department_response.Manager = &UserResponse{
			Id:        manager.Id,
			FullName:  manager.FullName,
			Email:     manager.Email,
			BirthDate: manager.DateOfBirth,
			Address:   manager.Address,
			NIK:       manager.NIK,
			Gender:    manager.Gender,
			Photo:     manager.Photo,
			Salary: SalaryResponse{
				Initial: manager_salary.Initial,
				Current: manager_salary.Current,
			},
			PhoneNumber: manager.PhoneNumber,
			FirstLogin:  manager.FirstLogin,
		}
	}

	manager_id := 0

	if department_head.ManagerId != nil {
		manager_id = *department_head.ManagerId
	}

	// Fetch search employee data
	employees, err :=
		models.User{}.SearchDepartmentEmployee(query, department.Id, manager_id)

	if err != nil {
		log.Panic("Error get employees: ", err.Error())

		return err
	}

	var user_responses []UserResponse

	for _, employee := range employees {
		var user_response UserResponse

		salary, err :=
			models.Salary{}.GetUsingUserId(employee.Id)

		// Ensure no error get employee salary
		if err != nil {
			log.Panic("Error get employee salary: ", err.Error())

			return err
		}

		user_response.Id = employee.Id
		user_response.FullName = employee.FullName
		user_response.Email = employee.Email
		user_response.Address = employee.Address
		user_response.BirthDate = employee.DateOfBirth
		user_response.PhoneNumber = employee.PhoneNumber
		user_response.Gender = employee.Gender
		user_response.NIK = employee.NIK
		user_response.Photo = employee.Photo
		user_response.Salary = SalaryResponse{
			Initial: salary.Initial,
			Current: salary.Current,
		}
		user_response.FirstLogin = employee.FirstLogin

		user_responses = append(user_responses, user_response)
	}

	department_response.Employees = &user_responses

	return nil
}
