package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load(".env") // Local env only
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}