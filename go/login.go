package forum

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginhandler called (go)")
}

// perform user lookup in user database
func authenticateUser() {
	fmt.Println("authenticateUser called")
}
