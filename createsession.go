package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSession(userName string, sessionToken string) {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	createSessionCacheTable(forumDatabase)
	insertNewSession(forumDatabase, userName, sessionToken)
}
func createSessionCacheTable(db *sql.DB) {
	createSessionSQL := `
		CREATE TABLE IF NOT EXISTS session(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT UNIQUE,
		"session-token" TEXT UNIQUE
	 );
	 `
	statement, err := db.Prepare(createSessionSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}
func insertNewSession(db *sql.DB, userName string, sessionToken string) {
	insertNewSessionSQL := `INSERT INTO session (username, session_token) VALUES (?, ?)`
	statement, err := db.Prepare(insertNewSessionSQL)
	//Prepares statement to avoid SQL injection
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(userName, sessionToken)
	if err != nil {
		log.Println(err.Error())
	}
}
