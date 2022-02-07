package forum

import (
	"database/sql"
	"log"

)

func LoginUser(userData []string, sessionToken string) {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	loginSQL := `SELECT username, password FROM user WHERE username = ? AND password = ?`
	row := forumDatabase.QueryRow(loginSQL, userData[0], userData[1])
	var userName string
	var password string
	err = row.Scan(&userName, &password)
	switch err {
	case sql.ErrNoRows:
		log.Println("Username or password incorrect!")
	case nil:
		log.Println("Logged in succesfully: ", userName, password)
		CreateSession(userName, sessionToken)
	default:
		log.Println(err.Error())
	}

	//displayUsers(forumDatabase)
}
