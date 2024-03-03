package forum

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginhandler called (go)")

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

	// Check if the user is already logged in
	var currentSessionID string
	err = DB.QueryRow("SELECT CurrentSessionID FROM Users WHERE Username = ?", storedUser.Username).Scan(&currentSessionID)
	if err != nil {
		fmt.Println("Error: user is already logged in")
	}

	if currentSessionID != "" {
		http.Error(w, "User is already logged in on a different browser", http.StatusForbidden)
		return
	}

	// Generate a new session ID
	newSessionID := generateSessionID(10) // Implement this function to generate a unique session ID

	// this updating the user's CurrentSessionID in the database
	_, err = DB.Exec("UPDATE Users SET CurrentSessionID = ? WHERE Username = ?", newSessionID, storedUser.Username)
	if err != nil {
		fmt.Println("Error: session id cannot be updated")
	}

	// Set the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: newSessionID,
		Path:  "/",
	})

	// Send a successful login response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s user logged in successfully\n", storedUser.Username)))
}

// perform user lookup in user database //below not implemented??
func authenticateUser() {
	fmt.Println("authenticateUser called")
}

// andle logging out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the session data from the database
	sessionCookie, err := r.Cookie("session")
	if err == nil {
		sessionID := strings.Split(sessionCookie.Value, "&")[0]
		clearSessionFromDB(sessionID)
	}

	// Clear session and user cookies
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, // Expire immediately
		Path:   "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		Value:  "",
		MaxAge: -1, // Expire immediately
		Path:   "/",
	})

	fmt.Println("user logged out successfully")

	// Redirect the user to the login page
	// http.Redirect(w, r, "/login", http.StatusFound)
}

// clearSessionFromDB removes the session data from the database
func clearSessionFromDB(sessionID string) {
	fmt.Println("clear session function called")
	// Implement code to clear the session from the database by setting SessionID to NULL
	query := "UPDATE Users SET SessionID = NULL WHERE SessionID = ?"
	_, err := DB.Exec(query, sessionID)
	if err != nil {
		// Handle the error, e.g., log it
		log.Println("Error clearing session from the database:", err)
	}
}
