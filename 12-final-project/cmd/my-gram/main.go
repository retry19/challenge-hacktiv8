package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	mygram "github.com/retry19/challenge-hacktiv8/12-final-project/api/my-gram"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
)

func main() {
	database.Init()
	defer database.Close()

	db := database.GetDB()
	app := fiber.New()

	server := mygram.NewServer(app, db)
	defer server.Close()

	go server.Start()

	sigTerm := make(chan os.Signal, 2)
	go func() {
		signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)
	}()
	<-sigTerm
}
