package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/retry19/challenge-hacktiv8/09-gemini-ai/controllers"
	"github.com/retry19/challenge-hacktiv8/09-gemini-ai/helpers"
)

func main() {
	helpers.LoadEnv()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", controllers.ViewAskPage)

	api := app.Group("/api")
	api.Post("/ask", controllers.AskHandler)

	app.Listen(":" + helpers.Port)
}
