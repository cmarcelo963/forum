package forum

import (
	"database/sql"
	"log"
)

func createPostTable(db *sql.DB) {
	createPostTableSQL := `
		CREATE TABLE IF NOT EXISTS post(
		"post_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT UNIQUE,
		"body" TEXT,
		"username" TEXT,
		"created_date" DATETIME,
		"category" TEXT
	 );
	 `
	statement, err := db.Prepare(createPostTableSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}
