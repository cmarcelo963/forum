package forum

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

//Store specific information from that post that was received
func HandleNewCommentRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postTitle := r.Form["title"][0]
	log.Println("title: ", postTitle)
	newCommentContent := r.Form["content"][0]
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	createCommentTable(forumDatabase)
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

	postID := getPostID(forumDatabase, postTitle)
	insertNewComment(forumDatabase, postID, username, newCommentContent)
	tpl, _ = template.ParseFiles("../static/templates/index.gohtml")
	// UserSession.SelectedPost = GetPost("", username)
	UserSession.Comments = GetComments(postID)
	tpl.Execute(w, UserSession)
}

//Adds relevant information of the new post into the database
func insertNewComment(db *sql.DB, post_id string, username string, content string) {
	insertNewCommentSQL := `INSERT INTO comment (post_id, content, username, created_date) VALUES (?, ?, ?, CURRENT_TIMESTAMP)`
	statement, err := db.Prepare(insertNewCommentSQL)

	if err != nil {
		log.Println(err.Error())
	}
	_, err = statement.Exec(post_id, content, username)
	if err != nil {
		log.Println(err.Error())
	}
}
func getPostID(db *sql.DB, title string) string {
	getPostIDSQL := `SELECT post_id FROM post WHERE title = ?`
	statement, err := db.Prepare(getPostIDSQL)
	if err != nil {
		log.Println(err.Error())
	}
	selectedPost, err := statement.Query(title)
	var postID string
	for selectedPost.Next() {
		err = selectedPost.Scan(&postID)
	}
	if err != nil {
		log.Println(err.Error())
	}
	return postID
}
