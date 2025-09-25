package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadDocs(t *testing.T) {
	// Create a temporary .adoc file
	_ = os.Mkdir("testdata", 0755)
	testContent := `
	[[test-section]]
	== Test Section

	This is a test section.
	`
	_ = os.WriteFile("testdata/test.adoc", []byte(testContent), 0644)
	defer os.RemoveAll("testdata")

	err := LoadDocs("testdata")
	if err != nil {
		t.Fatalf("failed to load docs: %v", err)
	}

	if len(Docs) != 1 {
		t.Fatalf("expected 1 doc entry, got %d", len(Docs))
	}

	if Docs[0].Title != "Test Section" {
		t.Errorf("expected title 'Test Section', got %s", Docs[0].Title)
	}

	if Docs[0].ID != "test-section" {
		t.Errorf("expected ID 'test-section', got %s", Docs[0].ID)
	}
}

func TestParseContent(t *testing.T) {
	Docs = nil
	DocContent = make(map[string]string)

	content := `[[test-section]]
	== Test Section
	This is a test section body.
	Another line.`

	filePath := "testdata/test.adoc"
	rootPath := "testdata"

	parseContent(filePath, rootPath, content)

	if len(Docs) != 1 {
		t.Fatalf("expected 1 doc entry, got %d", len(Docs))
	}

	doc := Docs[0]

	if doc.Title != "Test Section" {
		t.Errorf("expected title 'Test Section', got '%s'", doc.Title)
	}
	if doc.ID != "test-section" {
		t.Errorf("expected ID 'test-section', got '%s'", doc.ID)
	}

	expectedPath, _ := filepath.Rel(rootPath, filePath)
	if doc.File != expectedPath {
		t.Errorf("expected file path '%s', got '%s'", expectedPath, doc.File)
	}

	if doc.Snippet != "This is a test section body." {
		t.Errorf("expected snippet line mismatch: got '%s'", doc.Snippet)
	}

	if !strings.Contains(DocContent["test-section"], "Another line.") {
		t.Errorf("expected full content to include 'Another line.'")
	}
}
