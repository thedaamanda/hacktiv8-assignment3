package main

import (
	"assignment3/internal/jobs"
	"assignment3/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddress := setupServer()
	router := router.SetupRouter()

	jobs.StartScheduler()

	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("Error running server: " + err.Error())
	}
}

func setupServer() string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}
