package main

import (
	"fmt"
	"html/template"
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

	// files of templates
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	// func for read file of template home page
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Execute() for write template in body HTTP response
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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