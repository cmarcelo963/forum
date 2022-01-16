package forum

import "net/http"

func Server() {
	http.Handle("/index.css", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/", HandleDefault)
	http.ListenAndServe(":8080", nil)

}
