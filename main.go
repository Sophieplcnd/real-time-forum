package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	port := ":8080"
	http.ListenAndServe(port, nil)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postData map[string]string
	err := decoder.Decode(&postData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	title := postData["title"]
	content := postData["content"]

	db, err := sql.Open("sqlite3", ".data/database.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Post created successfully!"))
}
