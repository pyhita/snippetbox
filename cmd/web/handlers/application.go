package handlers

import (
	"github.com/pyhita/snippetbox/internal/models"
	"log"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Snippets *models.SnippetModel
}
