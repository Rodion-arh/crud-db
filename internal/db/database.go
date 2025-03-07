package db

import (
	"crud-db/internal/handler"
	"crud-db/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	var errConn error
	handler.DB, errConn = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if errConn != nil {
		log.Fatalf("Couldn't connect to the database: %v", errConn)
	}
	handler.DB.AutoMigrate(&models.Message{})

	fmt.Println("Database connected successfully!")
}
