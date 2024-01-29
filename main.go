package main

import (
	"fmt"
	"net/http"
	forum "rt-forum/go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	forum.Init()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

