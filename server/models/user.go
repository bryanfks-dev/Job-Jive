package models

import "time"

type User struct {
	Id           int       `json:"id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ManagerId    int       `json:"manager_id"`
	Address      string    `json:"address"`
	NIK          string    `json:"nik"`
	Gender       string    `json:"gender"`
	PhoneNumber  string    `json:"phone_number"`
	DepartmentId int       `json:"department_id"`
	FirstLogin   time.Time `json:"first_login"`
}

func (user User) PassChanged() {
	return 
}
