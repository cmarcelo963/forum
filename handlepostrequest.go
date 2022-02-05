package forum

import (
	"log"
	"net/http"
	// "text/template"
)

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) == 0 { // if the form contains no length or is a valid ascii character, 400 error
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	} else {
		email := r.Form["email"][0]
		//tpl, _ := template.ParseFiles("../templates/index.gohtml") // double check where the program is being run
		userName := r.Form["username"][0]
		password := r.Form["password"][0]
		var userData = []string{email, userName, password}
		log.Println(userData)
		//tpl.Execute(res, artToDisplay)
	}
}
