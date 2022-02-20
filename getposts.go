package forum

import (
	"database/sql"
	"log"
)

func GetPosts(category string) {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getPostSQL := `SELECT * FROM post WHERE categories LIKE ?`
	statement, err := forumDatabase.Prepare(getPostSQL)
	if err != nil {
		log.Println(err.Error())
	}
	filteredPosts, err := statement.Exec(category)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS >", filteredPosts.)
}
