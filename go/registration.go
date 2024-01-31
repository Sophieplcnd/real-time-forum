package forum

import (
	"fmt"
	"net/http"
)

type UserData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("go registerHandler triggered")
}
