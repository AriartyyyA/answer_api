package main

import (
	"fmt"
	"github/Ariartyyy/answer_api/internal/repository"
	"github/Ariartyyy/answer_api/internal/server"
	"github/Ariartyyy/answer_api/internal/service"
	"github/Ariartyyy/answer_api/internal/transport"
	"log"

	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	httpHandlers := transport.NewHTTPHandlers(service)
	httpServer := server.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}
	log.Println("server is starting")
}
