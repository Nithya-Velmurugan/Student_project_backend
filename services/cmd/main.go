package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"student-service/services/internal/handler"
	"student-service/services/internal/model"
	"student-service/services/internal/repository"
	"student-service/services/internal/service"
)

func main() {
	// Load .env file explicitly
	// Try loading from the current directory (if run from root)
	_ = godotenv.Load(".env")
	// Try loading from two directories up (if run from services/cmd)
	_ = godotenv.Load("../../.env")

	// Configure Database Connection
	// Default to a local postgres DB if not specified
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migration
	log.Println("Running Auto migrations...")
	err = db.AutoMigrate(
		&model.User{},
		&model.Student{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	// Add studentRepo as well if needed
	// studentRepo := repository.NewStudentRepository(db) // Assuming it exists

	// Initialize Services
	authService := service.NewAuthService(userRepo)
	// studentService := service.NewStudentService(studentRepo)

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(authService)

	// Initialize Gin Router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register Routes
	handler.RegisterAuthRoutes(r, authHandler)
	handler.RegisterStudentRoutes(r)

	// Add a ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
