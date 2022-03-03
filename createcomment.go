package forum

import (
	"database/sql"
	"log"
)

func createCommentTable(db *sql.DB) {
	createCommentTableSQL := `
		CREATE TABLE IF NOT EXISTS comment(
		"comment_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"post_id" integer,
		"content" TEXT,
		"username" TEXT,
		"created_date" DATETIME
	 );
	 `
	statement, err := db.Prepare(createCommentTableSQL)
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec()
}
