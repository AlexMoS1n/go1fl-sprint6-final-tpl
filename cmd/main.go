package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "LOG: ", log.LstdFlags)

	logger.Println("the server is running on port 8080")

	s:=server.CreateServerMorse(logger)

	err:=s.Server.ListenAndServe()
	if err != nil {
		logger.Fatalf("server startup error: %v", err)
	}
}
