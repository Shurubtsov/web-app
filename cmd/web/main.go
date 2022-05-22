package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"dshurubtsov.com/snippetbox/cmd/config"
)

// // struct for storage dependencies of logs and others
// type application struct {
// 	errorLog *log.Logger
// 	infoLog  *log.Logger
// }

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

	// main handlers for pages
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home(app))
	mux.HandleFunc("/snippet", ShowSnippet(app))
	mux.HandleFunc("/snippet/create", CreateSnippet(app))

	// initialize FileServer for proccesing HTTP requsts to static files from dir /ui/static
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	// using func mux.Handle() for registration proccesor for all requests which starts from "ui/static/"
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// initialize custom server for our error logs
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("start server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

// structure for defense static directory from client requests
type neuteredFileSystem struct {
	fs http.FileSystem
}

// method of our struct for check Open request
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
