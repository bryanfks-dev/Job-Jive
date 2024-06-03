package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	server_cmd := exec.Command("go", "run", "./server/main.go")
	server_cmd.Stdout = os.Stdout
	server_cmd.Stderr = os.Stderr

	// Ensure no error lauching server cmd
	if err := server_cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cron_job := exec.Command("go", "run", "./cron_job/main.go")
	cron_job.Stdout = os.Stdout
	cron_job.Stderr = os.Stderr

	// Ensure no error lauching cron_job cmd
	if err := cron_job.Start(); err != nil {
		log.Fatal(err)
	}

	// Ensure no error wait for cmd to execute
	if err := server_cmd.Wait(); err != nil {
		log.Println("Server process exited with error:", err)
	}

	// Ensure no error wait for cmd to execute
	if err := cron_job.Wait(); err != nil {
		log.Println("Cron job process exited with error:", err)
	}
}
