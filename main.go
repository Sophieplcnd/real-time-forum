package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     http.ServeFile(w, r, "index.html")
    // })

	// fmt.Println("Server listening on :8080...")
	// http.ListenAndServe(":8080", nil)
}

// func serveHTML(w http.ResponseWriter, r *http.Request) {
// 	content, err := os.ReadFile("index.html")
// 	if err != nil {
// 		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/html")
// 	w.Write(content)
// }
