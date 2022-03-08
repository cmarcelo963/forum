package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

type Auth struct {
	Authenticated     string
	AuthenticatedHide string
	User              string
	Posts             []Post
	SelectedPost      Post
	Comments          []Comment
}

var UserSession Auth

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			UserSession.Authenticated = ""
			UserSession.AuthenticatedHide = ""
			tpl.Execute(w, UserSession)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if AuthenticateSession(c.Value) {
		UserSession.Authenticated = "authenticated"
		UserSession.AuthenticatedHide = "authenticatedhide"
	} else {
		UserSession.Authenticated = ""
		UserSession.AuthenticatedHide = ""
	}
	tpl.Execute(w, UserSession)
}
