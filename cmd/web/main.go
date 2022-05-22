package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"dshurubtsov.com/snippetbox/cmd/config"
)

func main() {
	// configuration
	addr := flag.String("addr", ":4000", "Net address HTTP")
	flag.Parse()

	// configure new loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  // info
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // errors

	// init application object
	app := &config.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	// initialize custom server for our error logs
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  Routes(app),
	}

	infoLog.Printf("start server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
