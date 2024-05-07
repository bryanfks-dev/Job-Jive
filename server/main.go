package main

import (
	"fmt"
	"net/http"

	"configs"
	"db"
	"forms"

	"github.com/joho/godotenv"
)

func loadConfig() (configs.Server, configs.Database) {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	configs.Session.Init(configs.Session{})

	return configs.Server.Get(configs.Server{}), 
		configs.Database.Get(configs.Database{})
}

func initRoutes() {
	http.HandleFunc("/loginAuth", forms.LoginUserAuthHandler)
}

func main() {
	serverConf, dbConf := loadConfig()

	// Connect to database
	db.Connect(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	initRoutes()

	fmt.Printf("API Server is running on http://%s:%s\n", serverConf.Host, serverConf.Port)

	// Open server connection
	err := http.ListenAndServe(serverConf.Host + ":" + serverConf.Port, nil)

	if err != nil {
		panic(err.Error())
	}

	defer db.Conn.Close()
}