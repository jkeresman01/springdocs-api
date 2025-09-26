package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jkeresman01/springdocs-api/models"
)

var Docs []models.DocEntry
var DocContent = make(map[string]string)

func LoadDocs(rootPath string) error {
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".adoc" {
			return processFile(path, rootPath)
		}
		return nil
	})
}

func processFile(path, rootPath string) error {
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(contentBytes)
	parseContent(path, rootPath, content)
	return nil
}

func parseContent(filePath, rootPath, content string) {
	lines := strings.Split(content, "\n")
	var currentID, currentTitle string

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if isIDLine(line) {
			currentID = strings.Trim(line, "[]")
		} else if isTitleLine(line) {
			currentTitle = strings.TrimPrefix(line, "== ")
			if currentID != "" && currentTitle != "" {
				addDocEntry(currentTitle, currentID, filePath, rootPath, lines, i+1, content)
				currentID = ""
			}
		}
	}
}

func isIDLine(line string) bool {
	return strings.HasPrefix(line, "[[") && strings.Contains(line, "]]")
}

func isTitleLine(line string) bool {
	return strings.HasPrefix(line, "==")
}

func addDocEntry(title, id, fullPath, rootPath string, lines []string, snippetIndex int, content string) {
	snippet := ""
	if snippetIndex < len(lines) {
		snippet = strings.TrimSpace(lines[snippetIndex])
	}
	relPath, _ := filepath.Rel(rootPath, fullPath)

	entry := models.DocEntry{
		Title:   title,
		ID:      id,
		File:    relPath,
		Snippet: snippet,
	}

	Docs = append(Docs, entry)
	DocContent[id] = content
}
