package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	// Initialize
	err := timeInit()
	if err != nil {
		log.Fatalf("failed to init. %s\n", err)
	}

	// Connect to MQTT
	err = mqttInit()
	if err != nil {
		log.Fatalf("failed to connect to mqtt: %s\n", err)
	}

	// Run HTTP server
	runHttp()

	// Wait for exit
	ch := make(chan bool)
	<-ch
}
