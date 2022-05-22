package main

import (
	"net/http"
	"path/filepath"

	"dshurubtsov.com/snippetbox/cmd/config"
)

func Routes(app *config.Application) *http.ServeMux {

	// main routes for pages
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home(app))
	mux.HandleFunc("/snippet", ShowSnippet(app))
	mux.HandleFunc("/snippet/create", CreateSnippet(app))

	// initialize FileServer for proccesing HTTP requsts to static files from dir /ui/static
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	// using func mux.Handle() for registration proccesor for all requests which starts from "ui/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/static", http.NotFoundHandler())

	return mux
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
