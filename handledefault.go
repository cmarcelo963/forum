package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(w, nil)
}
