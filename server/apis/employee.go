package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

type Employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var employees = []Employee{}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range employees {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)
			employee.ID = params["id"]
			employees = append(employees, employee)
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}
