package forum

import (
	"database/sql"
	"log"
)

func getLastSessionToken(c string) {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	lastActiveSession := `SELECT * FROM session_cache WHERE session_token=(SELECT max(session_token) FROM session_cache)`
	log.Println(lastActiveSession)
}
