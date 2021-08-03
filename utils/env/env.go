package env

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("\u001b[31mError loading .env file!\u001b[0m")
		panic(err)
	}
}
