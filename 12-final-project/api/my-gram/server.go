package mygram

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

func NewServer(app *fiber.App, db *gorm.DB) *Server {
	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Start() error {
	s.app = fiber.New(fiber.Config{
		ServerHeader: "challenge-hacktiv8/12-final-project",
	})

	s.buildRoutes()

	return s.app.Listen(fmt.Sprintf(":%s", config.Port))
}

func (s *Server) Close() error {
	err := s.app.Shutdown()
	if err != nil {
		log.Error("failed to shutdown server", err.Error())
		return err
	}

	log.Info("server closed")

	return nil
}

func (s *Server) buildRoutes() {
	buildV1Routes(s.app, s.db)
}
