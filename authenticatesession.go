package forum

import (
	"database/sql"
	"log"
	"strings"
)

//Checks whether the user is currently logged in somewhere else to prevent double log-in
func AuthenticateSession(cookieValue string) bool {
	splitCookie := strings.SplitN(cookieValue, "-", 2)

	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	loginSQL := `SELECT MAX(id),username, session_token FROM session_cache WHERE username = ? `
	row := forumDatabase.QueryRow(loginSQL, splitCookie[0])
	log.Println("authenticate row: ", row)
	var sessionToken string
	var id int
	var username string
	row.Scan(&id, &username, &sessionToken)
	checkPwdMatch := splitCookie[1] == sessionToken
	log.Println(checkPwdMatch)
	if checkPwdMatch {
		return true
	} else {
		return false
	}
}
