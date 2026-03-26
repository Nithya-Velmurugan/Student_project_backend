package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"student-service/services/config"
	"student-service/services/internal/handler"
	"student-service/services/internal/repository"
	"student-service/services/internal/service"
)

func main() {
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