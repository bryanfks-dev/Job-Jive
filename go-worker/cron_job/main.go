package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
)

func main() {
	cron := cron.New()
	err := cron.AddFunc("*/1 * * * *", func() {
		fmt.Println("Wokrs")
	})

	if err != nil {
		panic(err.Error())
	}

	log.Println("Start cron job")
	cron.Start()

	select{}
}