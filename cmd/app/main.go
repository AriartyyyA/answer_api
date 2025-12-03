package main

import (
	"fmt"
	"github/Ariartyyy/answer_api/internal/repository"
	"github/Ariartyyy/answer_api/internal/server"
	"github/Ariartyyy/answer_api/internal/service"
	"github/Ariartyyy/answer_api/internal/transport"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var db *gorm.DB
	var err error

	// Retry connection logic for Docker
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database, retrying in 2s... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	httpHandlers := transport.NewHTTPHandlers(service)
	httpServer := server.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}
	log.Println("server is starting")
}
