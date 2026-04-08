package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "go")
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

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "The snippet id was : %v", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created new snippet...."))
}
