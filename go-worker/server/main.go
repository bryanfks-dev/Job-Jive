package main

import (
	"fmt"
	"net/http"

	"apis"
	"auths"
	"configs"
	"db"

	"github.com/joho/godotenv"
)

var (
	mux = http.NewServeMux()
)

func loadConfig() (configs.Server, configs.Database, error) {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		return configs.Server{}, configs.Database{}, err
	}

	return configs.Server.Get(configs.Server{}),
		configs.Database.Get(configs.Database{}), nil
}

func initEndPoints() {
	// Forms endpoints
	mux.HandleFunc("/auth/verify-token", auths.VerifyLoginToken)
	mux.HandleFunc("/auth/user/login", auths.UserLoginHandler)
	mux.HandleFunc("/auth/admin/login", auths.AdminLoginHandler)

	// API endpoints
	mux.HandleFunc("/api/user/profile", apis.GetUserProfileHandler)

	// Employees endpoints
	mux.HandleFunc("/api/users", apis.GetUsersHandler)

	// Department endpoints
	mux.HandleFunc("/api/departments", apis.GetDepartmentsHandler)
	mux.HandleFunc("/api/department/create", apis.CreateDepartmentHandler)
	mux.HandleFunc("/api/department/update/{id}", apis.UpdateDepartmentHandler)
	mux.HandleFunc("/api/department/delete/{id}", apis.DeleteDepartmentHandler)

	// Config endpoints
	mux.HandleFunc("/api/configs", apis.GetConfigsHandler)
	mux.HandleFunc("/api/configs/save", apis.SaveConfigsHandler)
}

func main() {
	serverConf, dbConf, err := loadConfig()

	// Ensure no error fetching config
	if err != nil {
		panic(err.Error())
	}

	// Connect to database
	err = db.Connect(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	// Ensure no error connecting to database
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Connected to database %s:%s (%s)\n", dbConf.Host, dbConf.Port, dbConf.Database)

	initEndPoints()

	fmt.Printf("API Server is running on http://%s:%s\n", serverConf.Host, serverConf.Port)
	fmt.Println("Logs:")

	// Open server connection
	err = http.ListenAndServe(serverConf.Host+":"+serverConf.Port, mux)

	if err != nil {
		panic(err.Error())
	}

	defer db.Conn.Close()
}
