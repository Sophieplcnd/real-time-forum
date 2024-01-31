// handle database insertions
package forum

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database initialised")
	defer DB.Close()
}

// func ServeLogin(w http.ResponseWriter, r *http.Request) {
// 	// Simulate content retrieval from a database or other source
// 	content := `
// 	<h2>Login</h2>
// 	<form id="login-form">
// 		<label for="username">Username:</label>
// 		<input type="text" id="username" name="username" required><br>

// 		<label for="password">Password:</label>
// 		<input type="password" id="password" name="password" required><br>

// 		<button type="button" id="login-button">Login</button>
// 	</form>
// 	`
// 	w.Header().Set("Content-Type", "text/html")
// 	w.Write([]byte(content))
// }
