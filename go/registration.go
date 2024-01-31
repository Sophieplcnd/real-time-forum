package forum

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
}

// Ensure your insertUser function handles the database connection properly.
func insertUser(user UserData) error {
	// Insert user data into the 'users' table (modify the query based on your table structure)
	_, err := DB.Exec("INSERT INTO Users (username, email, password, firstname) VALUES (?, ?, ?, ?)",
		user.Username, user.Email, user.Password, user.FirstName)
	if err != nil {
		return err
	}
	fmt.Println("insertUser function triggered")
	return nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("go registerHandler triggered")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var user UserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("user info decoded")

	// 500 error, cannot insert user data
	// Insert user data into the database
	err = insertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("user info inserted into database")

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("user registration successful (go)")
}
