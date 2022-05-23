package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"dshurubtsov.com/snippetbox/cmd/config"
	"dshurubtsov.com/snippetbox/pkg/models"
)

// handler for processing homepage
func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			app.ServerError(w, err)
			return
		}

		// Execute() for write template in body HTTP response
		err = ts.Execute(w, nil)
		if err != nil {
			app.ServerError(w, err)
		}
	}
}

// handler for showing snippet
func ShowSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// if query id is incorrect then we return not found error
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			app.NotFound(w)
			return
		}

		s, err := app.Snippets.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.NotFound(w)
			} else {
				app.ServerError(w, err)
			}
			return
		}

		// response
		fmt.Fprintf(w, "%v", s)
	}
}

//handler for create snippet
func CreateSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// settings for forbiddence requests without POST method
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.ClientError(w, http.StatusMethodNotAllowed)
			return
		}

		// test data
		title := "The story of the snail"
		content := "The snail crawled out of the shell,\npulled out its horns,\nand picked them up again."
		expires := "7"

		// transfer data to method app.Snippets.Insert() for create snippet and return ID
		id, err := app.Snippets.Insert(title, content, expires)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		// redirect user to relevant page of snippet
		http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	}
}
