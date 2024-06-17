package main

import (
	"fmt"
	"models"
	"net/http"

	"apis"
	"auths"
	"configs"
	"db"
)

var (
	mux           = http.NewServeMux()
	server_config = configs.Server{}
	db_config     = configs.Database{}
	jwt_config    = configs.JWT{}
)

func initEndPoints() {
	// Login endpoints
	mux.Handle("/auth/verify-token",
		auths.AuthenticationMiddlware(http.HandlerFunc(auths.VerifyToken)))
	mux.Handle("/auth/user/login",
		auths.AuthenticationMiddlware(http.HandlerFunc(auths.UserLoginHandler)))
	mux.Handle("/auth/admin/login",
		auths.AuthenticationMiddlware(http.HandlerFunc(auths.AdminLoginHandler)))

	// User
	mux.Handle("/api/user/profile",
		auths.AuthenticationMiddlware(
			auths.UserMiddleware(http.HandlerFunc(apis.GetUserProfileHandler))))
	mux.Handle("/api/user/attend",
		auths.AuthenticationMiddlware(
			auths.UserMiddleware(http.HandlerFunc(apis.AttendUserHandler))))
	mux.Handle("/api/user/attendance",
		auths.AuthenticationMiddlware(
			auths.UserMiddleware(http.HandlerFunc(apis.GetUserAttendanceHandler))))
	mux.Handle("/api/user/attendance/today/latest",
		auths.AuthenticationMiddlware(
			auths.UserMiddleware(http.HandlerFunc(apis.GetUserTodayLatestAttendanceHandler))))

	// Employees endpoints
	mux.Handle("/api/users",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.GetUsersHandler))))
	mux.Handle("/api/user/create",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.CreateUserHandler))))
	mux.Handle("/api/user/update/{id}",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.UpdateUserHandler))))
	mux.Handle("/api/user/delete/{id}",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.DeleteUserHandler))))

	// Department endpoints
	mux.Handle("/api/departments",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.GetDepartmentsHandler))))
	mux.Handle("/api/department/create",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.CreateDepartmentHandler))))
	mux.Handle("/api/department/update/{id}",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.UpdateDepartmentHandler))))
	mux.Handle("/api/department/delete/{id}",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.DeleteDepartmentHandler))))

	// Config endpoints
	mux.Handle("/api/configs",
		auths.AuthenticationMiddlware(http.HandlerFunc(apis.GetConfigsHandler)))
	mux.Handle("/api/configs/save",
		auths.AuthenticationMiddlware(
			auths.AdminMiddleware(http.HandlerFunc(apis.SaveConfigsHandler))))
}

func loadConfig(config configs.ConfigInterfaces) {
	err := config.Load()

	// Ensure no error load config
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	loadConfig(&server_config)
	loadConfig(&db_config)
	loadConfig(&jwt_config)

	// Assign secret key
	models.JWT_Secret = []byte(jwt_config.Secret)

	// Connect to database
	err :=
		db.Connect(db_config.User, db_config.Password,
			db_config.Host, db_config.Port, db_config.Database)

	// Ensure no error connecting to database
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Connected to database %s:%s (%s)\n", db_config.Host, db_config.Port, db_config.Database)

	initEndPoints()

	fmt.Printf("API Server is running on http://%s:%s\n", server_config.Host, server_config.Port)
	fmt.Println("Logs:")

	// Open server connection
	err = http.ListenAndServe(server_config.Host+":"+server_config.Port, mux)

	// Ensure no error
	if err != nil {
		panic(err.Error())
	}

	defer db.Conn.Close()
}
