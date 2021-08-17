package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	controllerV1 "github.com/rspindola/boilerplate_golang_api/controllers/controllers/api/v1"
	"github.com/rspindola/boilerplate_golang_api/seed"
)

var server = controllerV1.Server{}

func main() {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_CONNECTION"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))

	seed.Load(server.DB)

	server.Run(":8080")

	// cmd.Execute()

}
