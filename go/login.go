package forum

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginhandler called (go)")

	// below code is never reached
	var user UserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%s", user.Username)

	// Retrieve the user from the database
	var storedUser LoginData
	row := DB.QueryRow("SELECT Username, Email, PasswordHash FROM Users WHERE Username = ? OR Email = ?", user.Username, user.Email)
	err = row.Scan(&storedUser.Username, &storedUser.Email, &storedUser.Password)
	if err != nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	fmt.Printf("%s", storedUser.Username)

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	fmt.Printf("%s user logged in successfully\n", storedUser.Username)

}

// perform user lookup in user database
func authenticateUser() {
	fmt.Println("authenticateUser called")
}
