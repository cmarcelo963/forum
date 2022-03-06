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
func HandleNewLikeRequest(w http.ResponseWriter, r *http.Request) {
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
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	defer forumDatabase.Close()
	r.ParseForm()
	createLikeTable(forumDatabase)
	if err != nil {
		log.Println(err.Error())
	}
	if len(r.Form["post_id"]) > 0 {
		postID := r.Form["post_id"][0]
		insertNewPostLike(forumDatabase, postID, username, r.URL.Path)
	} else if len(r.Form["comment_id"]) > 0 {
		commentID := r.Form["comment_id"][0]
		insertNewCommentLike(forumDatabase, commentID, username, r.URL.Path)
	}
	
	tpl, _ = template.ParseFiles("../static/templates/index.gohtml")
	// UserSession.SelectedPost = GetPost("", username)
	// tpl.Execute(w, postSuccess)
}

//Adds relevant information of the new post into the database
func insertNewPostLike(db *sql.DB, post_id string, username string, like string) {
	if like == "/like" {
		like = "1"
	} else if like == "/dislike" {
		like = "0"
	}
	log.Println("HELLOOOOOOOOOOOOOOOOOOOOOO",like, "-",username, "-", post_id)
	updateNewLikeSQL := `UPDATE like SET like = ? WHERE username = ? AND post_id = ?`
	insertNewLikeSQL := `INSERT INTO like (post_id, comment_id, username, like) VALUES (?, null, ?, ?)`

	if !postLikeExists(db, username, post_id) {
		statement, err := db.Prepare(insertNewLikeSQL)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Inserting...")
		_, err = statement.Exec(post_id, username, like)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		statement, err := db.Prepare(updateNewLikeSQL)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Updating...")
		_, err = statement.Exec(like, username, post_id)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
func insertNewCommentLike(db *sql.DB, comment_id string, username string, like string) {
	if like == "/like" {
		like = "1"
	} else if like == "/dislike" {
		like = "0"
	}
	log.Println("HELLOOOOOOOOOOOOOOOOOOOOOO",like, "-",username, "-", comment_id)
	updateNewLikeSQL := `UPDATE like SET like = ? WHERE username = ? AND comment_id = ?`
	insertNewLikeSQL := `INSERT INTO like (post_id, comment_id, username, like) VALUES (null, ?, ?, ?)`

	if !commentLikeExists(db, username, comment_id) {
		statement, err := db.Prepare(insertNewLikeSQL)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Inserting...")
		_, err = statement.Exec(comment_id, username, like)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		statement, err := db.Prepare(updateNewLikeSQL)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Updating...")
		_, err = statement.Exec(like, username, comment_id)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
func postLikeExists(db * sql.DB, username string, post_id string) bool {
	var potatoarmy string
	checkIfLikeExistSQL := "SELECT * FROM like WHERE username = ? AND post_id = ?"
	err := db.QueryRow(checkIfLikeExistSQL, username, post_id).Scan(&potatoarmy)
	if err == sql.ErrNoRows {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		return false
	}
	return true
}
func commentLikeExists(db * sql.DB, username string, comment_id string) bool {
	var potatoarmy string
	checkIfLikeExistSQL := "SELECT * FROM like WHERE username = ? AND comment_id = ?"
	err := db.QueryRow(checkIfLikeExistSQL, username, comment_id).Scan(&potatoarmy)
	if err == sql.ErrNoRows {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		return false
	}
	return true
}