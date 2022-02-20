package forum

import (
	"database/sql"
	"log"
)

func LoginUser(userData []string, sessionToken string) bool {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	loginSQL := `SELECT username,password FROM user WHERE username = ?`
	row := forumDatabase.QueryRow(loginSQL, userData[0])
	log.Println(row)
	var userName string
	var password string
	row.Scan(&userName, &password)
	checkPwdMatch := ComparePasswords(password, []byte(userData[1]))
	log.Println("password: ", userData[1])
	log.Println("Hashed Password: ", password)

	//Checks whether the password the user inputted and the one saved in the database matches
	if checkPwdMatch {
		log.Println("Logged in succesfully: ", userName, password)
		CreateSession(userName, sessionToken)
		return true
	} else {
		log.Println("Username or password incorrect!")
		return false
	}

	// displayUsers(forumDatabase)
}
