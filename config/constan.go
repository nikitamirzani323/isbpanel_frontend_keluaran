package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const PATH_API string = "http://128.199.241.112:5051/"

func GetPathAPI() string {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}
	result := os.Getenv("PATH_API")
	log.Println(result)
	return result
}
