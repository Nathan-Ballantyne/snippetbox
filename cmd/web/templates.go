package main

import "github.com/Nathan-Ballantyne/snippetbox/pkg/models"

// Include a Snippets field in the templateData struct.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
