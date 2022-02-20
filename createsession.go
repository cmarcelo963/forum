package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Keeps track of user sessions
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
	//Creates a new table for session cache if one does not exist
	createSessionSQL := `
		CREATE TABLE IF NOT EXISTS session_cache(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT,
		"session_token" TEXT,
		"session_time" 	DATETIME
		);
	`
	statement, err := db.Prepare(createSessionSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}
func insertNewSession(db *sql.DB, userName string, sessionToken string) {
	insertNewSessionSQL := `INSERT INTO session_cache (username, session_token, session_time) VALUES (?, ?, CURRENT_TIMESTAMP)`
	statement, err := db.Prepare(insertNewSessionSQL)
	//timeStamp := `CURRENT_TIMESTAMP`
	//Prepares statement to avoid SQL injection
	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(userName, sessionToken)
	if err != nil {
		log.Println(err.Error())
	}
}
