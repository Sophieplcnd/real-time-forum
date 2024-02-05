// handle database
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
}

// close database
func CloseDatabase() {
	fmt.Println("CloseDatabase function called")
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Println("error closing database:", err)
		} else {
			fmt.Println("database closed successfully")
		}
	}
}
