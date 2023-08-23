package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Welcome to Go Weather App CLI")
	log.Println("This is a fun project create to help learn Go Lang and CLI tool")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error: failed to load env, please make sure you create valid .env file", err)
	}

	_, isPresent := os.LookupEnv("OPEN_WEATHER_API_KEY")
	if !isPresent {
		log.Fatalln("Cannot find OPEN_WEATHER_API_KEY variable in .env file, please make sure you have set the valid api key")
	}
}
