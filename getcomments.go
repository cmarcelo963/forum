package forum

import (
	"database/sql"
	"log"
)

type Comment struct {
	CommentId string
	Post_id	  string
	Content   string
	Username  string
	Date      string
}

func GetComments(post_id string) []Comment {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getCommentSQL := `SELECT * FROM comment WHERE post_id = ?`
	statement, err := forumDatabase.Prepare(getCommentSQL)
	if err != nil {
		log.Println(err.Error())
	}
	commentsRows, err := statement.Query(post_id)
	var comments []Comment
	for commentsRows.Next() {
		var c Comment
		err := commentsRows.Scan(&c.CommentId, &c.Post_id, &c.Content, &c.Username, &c.Date)
		if err != nil {
			log.Println(err.Error())
			break
		}
		log.Println("C > ", c)
		comments = append(comments, c)
	}
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS COMMENTS >>>>", comments, post_id)
	return comments
}
