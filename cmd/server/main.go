package main

import (
	"log"
	"simple-server/internal/server"
)

func main() {
	log.Println("Server start listening on port 8080")
	server.Start()
}