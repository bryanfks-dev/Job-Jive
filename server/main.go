package main

import (
	"fmt"
	"net/http"
	"os"

	"db"
	"forms"

	"github.com/joho/godotenv"
)

func initRoutes() {
	http.HandleFunc("/loginAuth", forms.LoginAuthHandler)
}

func main() {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	// Server credentials
	var (
		SERVER_URL = os.Getenv("SERVER_URL")
		SERVER_PORT = os.Getenv("SERVER_PORT")
	)

	// Connect to database
	db.Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), 
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	initRoutes()

	fmt.Printf("API Server is running on http://%s:%s\n", SERVER_URL, SERVER_PORT)

	// Open server connection
	err = http.ListenAndServe(SERVER_URL + ":" + SERVER_PORT, nil)

	if err != nil {
		panic(err.Error())
	}
}