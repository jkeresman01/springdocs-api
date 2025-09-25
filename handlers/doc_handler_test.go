package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jkeresman01/springdocs-api/models"
	"github.com/jkeresman01/springdocs-api/parser"
)

func setupTestApp() *fiber.App {
	app := fiber.New()

	parser.Docs = []models.DocEntry{{
		Title:   "Sample",
		ID:      "sample-id",
		File:    "sample.adoc",
		Snippet: "example snippet",
	}}
	parser.DocContent = map[string]string{
		"sample-id": "== Sample\nContent",
	}

	app.Get("/toc", GetTOC)
	app.Get("/search", SearchDocs)
	app.Get("/section", GetSection)

	return app
}

func TestGetTOC(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest("GET", "/toc", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestSearchDocs(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest("GET", "/search?q=sample", nil)
	resp, _ := app.Test(req)

	var result []map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&result)

	if len(result) != 1 || result[0]["id"] != "sample-id" {
		t.Fatalf("expected search to return sample-id")
	}
}

func TestGetSection(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest("GET", "/section?id=sample-id", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&result)

	if !strings.Contains(result["content"], "Sample") {
		t.Errorf("unexpected content: %v", result["content"])
	}
}
