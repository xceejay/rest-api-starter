// main.go

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xceejay/rest-api-starter/handlers"
)

const (
	ENVIRONMENT = "DEVELOPMENT"
)

func main() {

	server := handlers.Server{}
	log.Println("Configuration loaded:", ENVIRONMENT)
	server.Initialize(LoadDBConfig())
	log.Println("Starting server on http://127.0.0.1:7070")
	server.Run(":7070")
}

func LoadDBConfig() (string, string, string, string) {

	if ENVIRONMENT == "DEVELOPMENT" {
		err := godotenv.Load("config/.env.development")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		username := os.Getenv("USERNAME")
		password := os.Getenv("PASSWORD")
		hostname := os.Getenv("HOSTNAME")
		dbname := os.Getenv("DBNAME")
		return username, password, hostname, dbname

	} else {
		err := godotenv.Load("config/.env.production")
		if err != nil {
			log.Fatal("Error loading .env file")

		}
		username := os.Getenv("USERNAME")
		password := os.Getenv("PASSWORD")
		hostname := os.Getenv("HOSTNAME")
		dbname := os.Getenv("DBNAME")
		return username, password, hostname, dbname

	}

}
