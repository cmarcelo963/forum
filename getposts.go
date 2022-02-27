package forum

import (
	"database/sql"
	"log"
	"strings"
)

type Post struct {
	PostId          string
	Title           string
	Content         string
	Username        string
	Date            string
	Categories      string
	SplitCategories []string
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
		err := filteredPostsRows.Scan(&p.PostId, &p.Title, &p.Content, &p.Username, &p.Date, &p.Categories)
		if err != nil {
			log.Println(err.Error())
			break
		}
		splitCategories := strings.Split(p.Categories, ",")
		for index, category := range splitCategories {
			if index == len(splitCategories)-1 {
				break
			}
			p.SplitCategories = append(p.SplitCategories, category)
		}
		filteredPosts = append(filteredPosts, p)
	}
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS >", filteredPosts, category)
	return filteredPosts
}
