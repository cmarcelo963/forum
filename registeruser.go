package forum

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func RegisterUser(userDetails []string) {

	if _, err := os.Stat("./forum-database.db"); err != nil {
		file, err := os.Create("./forum-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}
	
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer forumDatabase.Close()

	createUserTable(forumDatabase)

	//INSERT RECORD OF NEW USER
	insertNewUser(forumDatabase, userDetails[0], userDetails[1], userDetails[2])

	displayUsers(forumDatabase)
}

func createUserTable(db *sql.DB) {
	createForumTableSQL := `
		CREATE TABLE IF NOT EXISTS user(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"email" TEXT,
		"username" TEXT,
		"password" TEXT
	 );
	 `

	statement, err := db.Prepare(createForumTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Print(err)
}

func insertNewUser(db *sql.DB, email string, username string, password string) {
	insertNewUserSQL := `INSERT INTO user (email, username, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNewUserSQL)
	//Prepares statement to avoid SQL injection

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(email, username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM user ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id uint
		var email string
		var username string
		var password string
		row.Scan(&id, &email, &username, &password)
		log.Printf("User: %v, %v, %v, %v", id, email, username, password)
	}
}
