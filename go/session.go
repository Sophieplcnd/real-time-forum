package forum

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
)

// we are generate a new session ID when a user logs in
// we are updating the user's CurrentSessionID in the database
// we are setting the session cookie
// we are checking if the user is already logged in on a different browser - still needs to be checked once login works
// we are querying the user's CurrentSessionID from the database
// if the user is already logged in, we are returning an error

// generateSessionID creates a new unique session ID
func generateSessionID(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// setSession creates a new session for the user and stores it in the database
func setSession(username string, w http.ResponseWriter, db *sql.DB) error {
	// Generate a new session ID
	newSessionID := generateSessionID(10)

	// Update the user's CurrentSessionID in the database
	_, err := db.Exec("UPDATE Users SET CurrentSessionID = ? WHERE Nickname = ?", newSessionID, username)
	if err != nil {
		return fmt.Errorf("could not set session in database: %v", err)
	}

	// Set the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: newSessionID,
		Path:  "/",
		// other cookie attributes...
	})

	return nil
}

// checkSession checks if the user is already logged in on a different browser
func checkSession(username string, db *sql.DB) (bool, error) {
	var currentSessionID string
	err := db.QueryRow("SELECT CurrentSessionID FROM Users WHERE Nickname = ?", username).Scan(&currentSessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found, so not logged in
			return false, nil
		}
		return false, fmt.Errorf("could not query current session: %v", err)
	}

	// User is logged in if currentSessionID is not empty
	return currentSessionID != "", nil
}
