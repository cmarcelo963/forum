package forum

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Auth struct {
	Authenticated string
}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
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

	var userSession Auth
	if AuthenticateSession(c.Value) {
		userSession.Authenticated = "authenticated"
	}
	log.Println(c, "hi")
	tpl.Execute(w, userSession)
	// getLastSessionToken(c.Value)
	// log.Println("Session token: ", c.Value)
}
