package forum

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func RegisterUser(userDetails []string) {
	if _, err := os.Stat("./forum-database.db"); err != nil {
		file, err := os.Create("./forum-database.db") // Create SQLite file
		if err != nil {
			log.Println(err.Error())
		}
		file.Close()
	}
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	encryptedPassword, _ := HashPassword(userDetails[2])
	createUserTable(forumDatabase)
	insertNewUser(forumDatabase, userDetails[0], userDetails[1], encryptedPassword)
	displayUsers(forumDatabase)
}

func createUserTable(db *sql.DB) {
	createForumTableSQL := `
		CREATE TABLE IF NOT EXISTS user(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"email" TEXT UNIQUE,
		"username" TEXT UNIQUE,
		"password" TEXT
	 );
	 `
	statement, err := db.Prepare(createForumTableSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}

func insertNewUser(db *sql.DB, email string, userName string, password string) {
	insertNewUserSQL := `INSERT INTO user (email, username, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNewUserSQL)
	// Prepares statement to avoid SQL injection
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(email, userName, password)
	if err != nil {
		log.Println(err.Error())
	}
}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM user ORDER BY id")
	if err != nil {
		log.Println(err)
	}
	defer row.Close()
	for row.Next() {
		var id uint
		var email string
		var username string
		var password string
		row.Scan(&id, &email, &username, &password)
		log.Printf("User: %v, %v, %v, %v", id, email, username, password)
	}
}
