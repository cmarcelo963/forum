package forum

import (
	"database/sql"
	"log"
)
type Post struct {
	PostId string
	Title string
	Content string
	Username string
	Date string
	Categories string
}
func GetPosts(category string) []Post {
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
	category = "%" + category + "%"
	filteredPostsRows, err := statement.Query(category)
	var filteredPosts []Post
	for filteredPostsRows.Next() {
		var p Post
		if err := filteredPostsRows.Scan(&p.PostId, &p.Title, &p.Content,&p.Username, &p.Date, &p.Categories); err != nil {
			log.Println(err.Error())
			break
		}
		filteredPosts = append(filteredPosts, p)
	}
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS >", filteredPosts, category)
	return filteredPosts
}
