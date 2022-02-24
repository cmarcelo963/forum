package forum

import (
	"log"
	"net/http"
	"text/template"
)
type FilteredPosts struct {
	Posts []Post
}
func HandleFilterRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) == 0 { // if the form contains no length or is a valid ascii character, 400 error
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	} else {
		log.Println("FORM: ", r.Form)
		category := r.Form["categories"][0]
		filteredPosts := GetPosts(category)
		UserSession.Posts = nil
		for _, post := range filteredPosts {
			UserSession.Posts = append(UserSession.Posts, post)
		}
		//&UserSession.Posts = GetPosts(category)
		tpl, _ := template.ParseFiles("../static/templates/login.gohtml")
		log.Println("kamal", UserSession)
		tpl.Execute(w, UserSession)
		//http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
	}
}
