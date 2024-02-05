package forum

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error encrypting password", err)
	}
	return hashedPassword
}

func insertUser(user UserData) error {
	hashedPassword := hashPassword(user.Password)

	fmt.Println("inserUser function called")
	_, err := DB.Exec("INSERT INTO Users (Username, Email, PasswordHash, FirstName, LastName, Age, Gender) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.Username, user.Email, hashedPassword, user.FirstName, user.LastName, user.Age, user.Gender)
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
