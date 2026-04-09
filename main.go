package main

import (
	"fmt"
	"log"
	"os"
)

// event-bus — In-process event bus with typed events and async handlers
func main() {
	logger := log.New(os.Stdout, "[event-bus] ", log.LstdFlags)
	logger.Println("Starting application...")

	if err := run(); err != nil {
		logger.Fatalf("Fatal error: %v", err)
	}
}

func run() error {
	fmt.Println("Application initialized successfully")
	return nil
}
