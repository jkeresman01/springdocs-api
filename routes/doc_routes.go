package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/springdocs-api/handlers"
)

func RegisterDocRoutes(app *fiber.App) {
	app.Get("/toc", handlers.GetTOC)
	app.Get("/search", handlers.SearchDocs)
	app.Get("/section", handlers.GetSection)
}
