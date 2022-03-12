package forum

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Auth struct {
	Authenticated     string
	AuthenticatedHide string
	User              string
	CreatedPosts      []Post
	LikedPosts        []Post
	Posts             []Post
	SelectedPost      Post
	Comments          []Comment
	Likes             string
}

var UserSession Auth

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	sessionState := CheckSessionCookie(w, r, tpl)
	if sessionState == "No cookie" || sessionState == "Bad request" {
		return
	}
	tpl.Execute(w, UserSession)
}
func CheckSessionCookie(w http.ResponseWriter, request *http.Request, template *template.Template) string {
	c, err := request.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			UserSession.Authenticated = ""
			UserSession.AuthenticatedHide = ""
			template.Execute(w, UserSession)
			return "No cookie"
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return "Bad request"
	}
	if AuthenticateSession(c.Value) {
		UserSession.Authenticated = "authenticated"
		UserSession.AuthenticatedHide = "authenticatedhide"
		return "Valid cookie"
	} else {
		log.Println("Not authenticated")
		UserSession.Authenticated = ""
		UserSession.AuthenticatedHide = ""
		UserSession.CreatedPosts = nil
		UserSession.LikedPosts = nil
		return "Invalid cookie"
	}
}


