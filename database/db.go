package database

import (
	"fmt"
	"log"
	"os"

	"web-service-gin/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file : %v", err)
	}

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error reaching database : %v", err)
	}

	DB = db
	fmt.Println("Successfully connected to database!")

	err = DB.AutoMigrate(&models.Album{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	fmt.Println("Database migration completed!")
}
