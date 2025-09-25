package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/springdocs-api/parser"
	"github.com/jkeresman01/springdocs-api/routes"
)

func main() {
	app := fiber.New()

	if err := parser.LoadDocs("./spring-docs"); err != nil {
		log.Fatalf("Failed to load docs: %v", err)
	}

	routes.RegisterDocRoutes(app)

	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
