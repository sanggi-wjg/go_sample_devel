package main

import (
	"github.com/joho/godotenv"
	"go_sample_devel/pkg/database"
	"go_sample_devel/route"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}
	envDatabase := &database.EnvDatabase{
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		DatabaseName: os.Getenv("DB_DATABASE_NAME"),
	}
	if err := database.Setup(envDatabase); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := route.SetupRoutes()
	log.Fatalln(router.Run(":9091"))
}
