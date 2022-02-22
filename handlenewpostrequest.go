package forum

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)
type PostSuccess struct {
	IsSuccessful bool
}
//Store specific information from that post that was received
func HandleNewPostRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newPostTitle := r.Form["title"][0]
	newPostContent := r.Form["content"][0]
	newPostCategories := ""
	for index, field := range r.Form {
		if field[0] == "on" {
			log.Println(field, index)
			newPostCategories += index + ","
			log.Println(newPostCategories)
		}
	}

	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	createPostTable(forumDatabase)
	//I changed the err into '_' to remove the orange error as it states that it isn't being used - cmarcelo963
	tpl, _ := template.ParseFiles("../static/templates/index.gohtml")
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			tpl.Execute(w, nil)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := strings.SplitN(c.Value, "-", 2)[0]

	var postSuccess PostSuccess
	postSuccess.IsSuccessful = insertNewPost(forumDatabase, username, newPostTitle, newPostContent, newPostCategories)
	tpl, _ = template.ParseFiles("../static/templates/create-post.gohtml")
	tpl.Execute(w, postSuccess)
}
//Adds relevant information of the new post into the database
func insertNewPost(db *sql.DB, username string, title string, content string, categories string) bool {
	insertNewPostSQL := `INSERT INTO post (username, title, created_date, content, categories) VALUES (?, ?, CURRENT_TIMESTAMP, ?, ?)`
	statement, err := db.Prepare(insertNewPostSQL)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = statement.Exec(username, title, content, categories)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
