package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"student-service/services/config"
	"student-service/services/internal/handler"
	"student-service/services/internal/repository"
	"student-service/services/internal/service"
)

func main() {
	// Attempt to load .env from various possible paths.
	// Using Overload ensures it overwrites any existing empty variables.
	envPaths := []string{".env", "../.env", "../../.env", "/home/ubuntu/Student_project_backend/.env"}
	var loaded bool
	for _, path := range envPaths {
		if err := godotenv.Overload(path); err == nil {
			log.Println("✅ Successfully loaded config from:", path)
			loaded = true
			break
		}
	}
	if !loaded {
		log.Println("⚠️ WARNING: Could not find or load a .env file. Using system environment variables.")
	}

	config.ConnectDB()

	if config.DB == nil {
		log.Fatal("Failed to connect to database in main")
	}

	userRepo := repository.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	// Reliable CORS config setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// routes
	handler.RegisterAuthRoutes(r, authHandler)
	handler.RegisterStudentRoutes(r)

	r.Run(":8082")
}