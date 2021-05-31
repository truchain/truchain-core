package main

import (
	"log"
	"truchain-core/networking"

	"github.com/joho/godotenv"
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

	resp, err := networking.IpApiGetIPLocation(ip)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Node Location: " + resp.City)
}
