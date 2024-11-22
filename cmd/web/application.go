package main

import (
	"log"

	"github.com/pyhita/snippetbox/internal/models"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Snippets *models.SnippetModel
}
