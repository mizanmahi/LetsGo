package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)



var AppConfig struct {
	Port string
}


func LoadEnvs () {
	err := godotenv.Load()
	if err  != nil {
		log.Println("No .env file found")
	}

	AppConfig.Port = os.Getenv("PORT")
	if AppConfig.Port == "" {
		AppConfig.Port = "5050"
	}
}