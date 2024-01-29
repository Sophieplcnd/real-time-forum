package main

import (
  "fmt"
	"database/sql"
	"net/http"
	forum "rt-forum/go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
   forum.Init()
  
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	port := ":8080"
	http.ListenAndServe(port, nil)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("post-title")
	content := r.FormValue("post-content")
}


