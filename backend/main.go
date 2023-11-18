package main

import (
	"log"
	server "subman/api/server"
	indexer "subman/indexer"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	go indexer.Start()
	server.Start()
}
