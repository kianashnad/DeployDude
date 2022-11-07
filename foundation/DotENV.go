package foundation

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadENVs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetENV(variableName string) string {
	envVariable := os.Getenv(variableName)
	if len(envVariable) == 0 {
		log.Fatalf("Error loading env variable: %v", envVariable)
	}

	return envVariable
}
