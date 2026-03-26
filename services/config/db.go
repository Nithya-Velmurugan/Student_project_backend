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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	var db *gorm.DB
var err error

for i := 0; i < 10; i++ {
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err == nil {
        log.Println("✅ Connected to DB")
        break
    }

    log.Println("⏳ Waiting for DB...")
    time.Sleep(2 * time.Second)
}

if err != nil {
    log.Fatal("❌ Failed to connect DB:", err)
}
}
