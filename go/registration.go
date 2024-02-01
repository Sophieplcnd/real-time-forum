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
	// LastName  string `json:"lastname"`
	// Age       int    `json:"age"`
}

func insertUser(user UserData) error {
	fmt.Println("inserUser function called")
	_, err := DB.Exec("INSERT INTO Users (Username, Email, PasswordHash, FirstName) VALUES (?, ?, ?, ?)",
		user.Username, user.Email, user.Password, user.FirstName)
	if err != nil {
		fmt.Println("error inserting user into database:", err)
		return err
	}
	fmt.Println("insertUser function successful")
	return nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registerHandler function called")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user UserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("user info decoded")

	err = insertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("user info inserted into database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("user registration successful (go)")
}
