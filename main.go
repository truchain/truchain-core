package main

import (
	"github.com/joho/godotenv"
	"log"
	"truchain-core/networking"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ip, err := networking.GetCurrentNodeIP()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Node IP: " + ip)

	resp, err := networking.GetIPLocation(ip)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Node Location: " + resp.City)
}
