package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"dshurubtsov.com/snippetbox/cmd/config"
	"dshurubtsov.com/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// configuration
	addr := flag.String("addr", ":4000", "Net address HTTP")
	dsn := flag.String("dsn", "web:0907@/snippetbox?parseTime=true", "name MySQL data source")
	flag.Parse()

	// configure new loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  // info
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // errors

	// configure database connection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// init application object
	app := &config.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Snippets: &mysql.SnippetModel{DB: db},
	}

	// initialize custom server for our error logs
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  Routes(app),
	}

	infoLog.Printf("start server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// sql connections pull
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
