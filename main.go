package main

import (
	"encoding/json"
	"log"
	"truchain-core/wallet"
)

func main() {
	response, err := wallet.GenerateWallet()
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.MarshalIndent(response, "", "  ")

	log.Print(string(data))
}
