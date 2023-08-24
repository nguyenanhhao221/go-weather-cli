package main

import (
	"bufio"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Welcome to Go Weather App CLI")
	log.Println("This is a fun project create to help learn Go Lang and CLI tool")

	// Handle ENV for the Open weather api key
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: failed to load env, please make sure you create valid .env file", err)
	}

	_, isPresent := os.LookupEnv("OPEN_WEATHER_API_KEY")
	if !isPresent {
		log.Fatalln("Cannot find OPEN_WEATHER_API_KEY variable in .env file, please make sure you have set the valid api key")
	}

	log.Print("Input your city: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cityInput := scanner.Text()
	log.Printf("Your input: %s", cityInput)
}
