package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler for processing homepage
func home(w http.ResponseWriter, r *http.Request) {
	// Check for correct input with right path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello world! homepage"))
}

// handler for showing snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {

	// if query id is incorrect then we return not found error
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Snippet with ID %d", id)
}

//handler for create snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {

	// settings for forbiddence requests without POST method
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Forbidden method!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create snippet here!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
