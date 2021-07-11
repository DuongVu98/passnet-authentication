package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	var err = godotenv.Load(getEnvFile())

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func getEnvFile() string {
	envFolder := "env/"
	env := os.Getenv("ENV")
	if env == "development" {
		return fmt.Sprintf("%v.env.dev", envFolder)
	}
	return ""
}
