package forum

import "net/http"

func Server() {
	http.Handle("/index.css", http.FileServer(http.Dir("../static")))
	http.Handle("/search.svg", http.FileServer(http.Dir("../static/img")))
	http.Handle("/calendar.svg", http.FileServer(http.Dir("../static/img")))
	http.Handle("/home.svg", http.FileServer(http.Dir("../static/img")))
	http.Handle("/about.svg", http.FileServer(http.Dir("../static/img")))
	http.Handle("/contact.svg", http.FileServer(http.Dir("../static/img")))
	http.Handle("/press.svg", http.FileServer(http.Dir("../static/img")))
	http.Handle("/random.svg", http.FileServer(http.Dir("../static/img")))
	http.HandleFunc("/", HandleDefault)
	http.ListenAndServe(":8080", nil)
}
