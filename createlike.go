package forum

import (
	"database/sql"
	"log"
)

func createLikeTable(db *sql.DB) {
	createLikeTableSQL := `
		CREATE TABLE IF NOT EXISTS like(
		"like_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"comment_id" integer,
		"post_id" integer,
		"username" TEXT,
	 );
	 `
	statement, err := db.Prepare(createLikeTableSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}
