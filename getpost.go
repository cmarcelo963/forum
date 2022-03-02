package forum

import (
	"database/sql"
	"log"
	"strings"
)

func GetPost(createdDate string, username string) Post {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getPostSQL := `SELECT * FROM post WHERE created_date = ? AND username = ?`
	statement, err := forumDatabase.Prepare(getPostSQL)
	if err != nil {
		log.Println(err.Error())
	}
	selectedPost, err := statement.Query(createdDate, username)
	var p Post
	err = selectedPost.Scan(&p.PostId, &p.Title, &p.Content, &p.Username, &p.Date, &p.Categories)
	if err != nil {
		log.Println(err.Error())
	}
	splitCategories := strings.Split(p.Categories, ",")
	for index, category := range splitCategories {
		if index == len(splitCategories)-1 {
			break
		}
		p.SplitCategories = append(p.SplitCategories, category)
	}
	log.Println("THIS selected >", createdDate, username, p.Title)
	return p
}
