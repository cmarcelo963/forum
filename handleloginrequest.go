package forum

import (
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
)
type LoginSuccess struct {
	IsSuccessful bool
}
func HandleLoginRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) == 0 { // if the form contains no length or is a valid ascii character, 400 error
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	} else {
		tpl, _ := template.ParseFiles("../static/templates/login.gohtml") // double check where the program is being run
		userName := r.Form["username"][0]
		password := r.Form["password"][0]
		var userData = []string{userName, password}
		sessionToken := uuid.NewV4().String()
		authenticated := LoginUser(userData, sessionToken)
		var loggedIn LoginSuccess
		loggedIn.IsSuccessful = false
		if authenticated {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   userName + "-" + sessionToken,
				Expires: time.Now().Add(30 * time.Minute),
			})
			loggedIn.IsSuccessful = true
		}
		tpl.Execute(w, loggedIn)
	}
}
