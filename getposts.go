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
	Likes           string
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
		p.Likes = GetLikes(p.PostId, "post")
		filteredPosts = append(filteredPosts, p)
	}
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS >", filteredPosts, category)
	return filteredPosts
}
func GetUserPosts(username string) []Post {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getUserPostsSQL := `SELECT * FROM post WHERE username = ?`
	statement, err := forumDatabase.Prepare(getUserPostsSQL)
	if err != nil {
		log.Println(err.Error())
	}
	filteredUserPostsRows, err := statement.Query(username)
	var userPosts []Post
	for filteredUserPostsRows.Next() {
		var p Post
		err := filteredUserPostsRows.Scan(&p.PostId, &p.Title, &p.Content, &p.Username, &p.Date, &p.Categories)
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
		p.Likes = GetLikes(p.PostId, "post")
		userPosts = append(userPosts, p)
	}
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS userposts>", userPosts, username)
	return userPosts
}
func GetUserLikedPosts(username string) []Post {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getLikedPostsSQL := `SELECT post.post_id, post.title, post.content, post.username, post.created_date, post.categories, like.like FROM post JOIN like ON post.post_id = like.post_id WHERE like.username = ?;`
	statement, err := forumDatabase.Prepare(getLikedPostsSQL)
	if err != nil {
		log.Println(err.Error())
	}
	likedPostsRows, err := statement.Query(username)
	var likedPosts []Post
	for likedPostsRows.Next() {
		var p Post
		err := likedPostsRows.Scan(&p.PostId, &p.Title, &p.Content, &p.Username, &p.Date, &p.Categories, &p.Likes)
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
		p.Likes = GetLikes(p.PostId, "post")
		likedPosts = append(likedPosts, p)
	}
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("THIS userposts>", likedPosts, username)
	return likedPosts
}