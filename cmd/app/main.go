package main

import (
	"fmt"
	"github/Ariartyyy/answer_api/internal/server"
	"github/Ariartyyy/answer_api/internal/transport"
	"log"
)

func main() {
	httpHandlers := transport.NewHTTPHandlers()
	httpServer := server.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}
	log.Println("server is starting")
}
