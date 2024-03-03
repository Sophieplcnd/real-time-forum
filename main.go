package main

import (
	"fmt"
	"net/http"

	forum "github.com/Sophieplcnd/real-time-forum/go"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	forum.Init()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/register", forum.RegisterHandler)
	http.HandleFunc("/login", forum.LoginHandler)
	http.HandleFunc("/logout", forum.LogoutHandler)

	port := ":8080"
	fmt.Println("Fetching server...")
	http.ListenAndServe(port, nil)

	// shutdown database when page is closed
	forum.CloseDatabase()
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// title := r.FormValue("post-title")
	// content := r.FormValue("post-content")
}
