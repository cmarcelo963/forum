package forum

import (
	"database/sql"
	"log"
)

func createLikeTable(db *sql.DB) {
	createLikeTableSQL := `
		CREATE TABLE IF NOT EXISTS like(
		"like_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"comment_id" INTEGER,
		"post_id" INTEGER,
		"username" TEXT,
		"like" BIT
	 );
	 `
	statement, err := db.Prepare(createLikeTableSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}
