package main

import (
	"fmt"
	"net/http"

	"configs"
	"db"
	"forms"

	"github.com/joho/godotenv"
)

var (
	mux = http.NewServeMux()
)

func loadConfig() (configs.Server, configs.Database) {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	return configs.Server.Get(configs.Server{}),
		configs.Database.Get(configs.Database{})
}

func initEndPoints() {
	mux.HandleFunc("/auth/user/login", forms.UserLoginHandler)
	mux.HandleFunc("/auth/admin/login", forms.AdminLoginHandler)
}

func main() {
	serverConf, dbConf := loadConfig()

	// Connect to database
	db.Connect(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	initEndPoints()

	fmt.Printf("API Server is running on http://%s:%s\n", serverConf.Host, serverConf.Port)
	fmt.Println("Logs:")

	// Open server connection
	err := http.ListenAndServe(serverConf.Host+":"+serverConf.Port, mux)

	if err != nil {
		panic(err.Error())
	}

	defer db.Conn.Close()
}
