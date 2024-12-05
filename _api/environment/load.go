package environment

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}
