package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippet box"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	response := fmt.Sprintf("The snippet : %v", id)
	w.Write([]byte(response))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display a form to create snippet...")
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "go")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "The snippet was saves successfully")
}

func main() {
	fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}/{$}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Println("The server is running on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
