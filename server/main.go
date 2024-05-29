package main

import (
	"fmt"
	"io"
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
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string("wllwlw"))
	})
	
	mux.HandleFunc("/loginAuth", forms.LoginUserAuthHandler)
}

func main() {
	serverConf, dbConf := loadConfig()

	// Connect to database
	db.Connect(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	initEndPoints()

	fmt.Printf("API Server is running on http://%s:%s\n", serverConf.Host, serverConf.Port)

	// Open server connection
	err := http.ListenAndServe(serverConf.Host + ":" + serverConf.Port, mux)

	if err != nil {
		panic(err.Error())
	}

	defer db.Conn.Close()
}