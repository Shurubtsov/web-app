package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// initialize FileServer for proccesing HTTP requsts to static files from dir /ui/static
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static/")})

	// using func mux.Handle() for registration proccesor for all requests which starts from "ui/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
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
