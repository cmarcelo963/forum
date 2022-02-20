package forum

import (
	"log"
	"net/http"
)

func HandleFilterRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) == 0 { // if the form contains no length or is a valid ascii character, 400 error
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	} else {
		log.Println("FORM: ", r.Form)
		category := r.Form["categories"][0]
		GetPosts(category)
		// tpl, _ := template.ParseFiles("../static/templates/index.gohtml")
		//tpl, _ := template.ParseFiles("../templates/index.gohtml") // double check where the program is being run

		// tpl.Execute(w, signedUp)
		//http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
	}
}
