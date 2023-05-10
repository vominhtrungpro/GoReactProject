package main

import (
	"GoReactProject/internal/handler/rest"
	"log"

	_ "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rest.ProjectAPI()
}
