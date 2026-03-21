package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"student-service/services/internal/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Try loading .env variables
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "4207"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "studentdb"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	var dsn string
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		dsn = dbURL
	} else if pgURL := os.Getenv("POSTGRES_URL"); pgURL != "" {
		dsn = pgURL
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, user, password, dbname, port, sslmode)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to postgres: %v\n", err)
	} else {
		DB = db
		fmt.Println("PostgreSQL connection established.")

		// Auto Migrate
		fmt.Println("Running AutoMigration...")
		err = DB.AutoMigrate(&model.User{}, &model.Student{})
		if err != nil {
			log.Printf("Failed to migrate database: %v\n", err)
		}
	}
}
