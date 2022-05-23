package config

import (
	"log"

	"dshurubtsov.com/snippetbox/pkg/models/mysql"
)

// struct for storage dependencies of logs and others
type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Snippets *mysql.SnippetModel
}
