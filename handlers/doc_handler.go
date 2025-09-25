package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/springdocs-api/parser"
)

func GetTOC(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(parser.Docs)
}

func SearchDocs(c *fiber.Ctx) error {
	query := strings.ToLower(c.Query("q"))
	var result []any

	for _, doc := range parser.Docs {
		if strings.Contains(strings.ToLower(doc.Title), query) ||
			strings.Contains(strings.ToLower(doc.Snippet), query) {
			result = append(result, doc)
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetSection(c *fiber.Ctx) error {
	id := c.Query("id")
	content, exists := parser.DocContent[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Section not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"content": content,
	})
}
