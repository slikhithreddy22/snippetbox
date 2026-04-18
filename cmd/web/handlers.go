package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	// "text/template"

	"github.com/slikhithreddy22/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// for _, snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v", snippet)
	// }
	files := []string{
		"ui/html/base.html",
		"ui/html/pages/home.html",
		"ui/html/partials/nav.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data := templateData{
		Snippets: snippets,
	}
	//
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	app.render(w, r, http.StatusOK, "home.tmpl", templateData{
		Snippets: snippets,
	})
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	app.render(w, r, http.StatusOK, "home.tmpl", templateData{
		Snippet: snippet,
	})
	files := []string{
		"ui/html/base.html",
		"ui/html/pages/view.html",
		"ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", snippet)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
