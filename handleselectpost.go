package forum

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func HandleSelectPost(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	sessionState := CheckSessionCookie(w, r, tpl)
	if sessionState == "no cookie" || sessionState == "bad request" {
		return
	}
	r.ParseForm()
	date := r.Form["date"][0]
	username := r.Form["username"][0]
	date = strings.Join(strings.Split(strings.Join(strings.Split(date, "T"), " "), "Z"), "")
	post := GetPost(date, username)
	likes := GetLikes(post.PostId, "post")
	comments := GetComments(post.PostId)
	UserSession.Comments = nil
	UserSession.Likes = likes
	for _, comment := range comments {
		UserSession.Comments = append(UserSession.Comments, comment)
	}
	log.Println("COMMENTS:", UserSession.Comments)
	UserSession.SelectedPost = post
	if err != nil {
		log.Println(err.Error())
	}
	tpl.Execute(w, UserSession)
}
func GetLikes(id string, typeOfMessage string) string {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getLikesSQL := `SELECT (SELECT COUNT(like) from like WHERE ` + typeOfMessage + `_id = ? and like = 1) - (SELECT COUNT(like) from like WHERE ` + typeOfMessage + `_id = ? AND like = 0);`
	statement, err := forumDatabase.Prepare(getLikesSQL)
	if err != nil {
		log.Println(err.Error())
	}
	var likes string
	err = statement.QueryRow(id, id).Scan(&likes)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Likes: ", likes)
	return likes
}
