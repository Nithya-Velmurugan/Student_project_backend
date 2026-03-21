package app

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

var Engine *gin.Engine

func InitApp() *gin.Engine {
	if Engine != nil {
		return Engine
	}

	// Try loading from the current directory or parent directory
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../../.env")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=4207 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Running Auto migrations...")
	err = db.AutoMigrate(
		&model.User{},
		&model.Student{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://18.61.174.8:5173",
			"http://18.61.174.8",
			"https://student-project-frontend-eight.vercel.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler.RegisterAuthRoutes(r, authHandler)
	handler.RegisterStudentRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	Engine = r
	return r
}
