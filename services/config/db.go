package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"student-service/services/internal/model"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL")
	sslrootcert := os.Getenv("DB_SSLROOTCERT")

	if host == "" || sslmode == "" {
		log.Fatal("🚨 CRITICAL ERROR: DB_HOST or DB_SSL is empty! The .env file was NOT loaded correctly.")
	}

	log.Println("HOST:", host)
	log.Println("SSL:", sslmode)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
		sslrootcert,
	)

	var err error

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to DB")
			
			// Auto Migrate
			fmt.Println("Running AutoMigration...")
			err = DB.AutoMigrate(&model.User{}, &model.Student{})
			if err != nil {
				log.Printf("Failed to migrate database: %v\n", err)
			}
			return
		}

		log.Println("⏳ Waiting for DB...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Failed to connect DB:", err)
}