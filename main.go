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
	timeInit()

	// Connect to MQTT
	err := mqttInit()
	if err != nil {
		log.Fatalln("failed to connect to mqtt")
	}

	// Run HTTP server
	runHttp()

	// Wait for exit
	ch := make(chan bool)
	<-ch
}
