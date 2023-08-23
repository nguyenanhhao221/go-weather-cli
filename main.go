package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Welcome to Go Weather App CLI")
	log.Println("This is a fun project create to help learn Go Lang and CLI tool")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error: failed to load env, please make sure you create valid .env file", err)
	}

	// openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")
}
