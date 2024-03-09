package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("no .env file, will use loaded environment values")
	}
}
