package main

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	postMu sync.Mutex
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	postMu.Lock()
	defer postMu.Unlock()

	r.Header.Set("Content-type", "application/json")

	// Logic here
	json.NewEncoder(w).Encode(/* This param should be a struct */ nil)
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

	// Server endpoints
	http.HandleFunc("/api/x", handleHTTP)

	fmt.Printf("API Server is running on http://%s:%s\n", SERVER_URL, SERVER_PORT)

	// Open server connection
	err = http.ListenAndServe(SERVER_URL + ":" + SERVER_PORT, nil)

	if err != nil {
		panic(err.Error())
	}
}