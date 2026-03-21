package main

import (
	"log"
	"os"
	"student-service/services/app"
)

func main() {
	r := app.InitApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
