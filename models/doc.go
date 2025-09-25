package models

type DocEntry struct {
	Title   string `json:"title"`
	ID      string `json:"id"`
	File    string `json:"file"`
	Snippet string `json:"snippet,omitempty"`
}
