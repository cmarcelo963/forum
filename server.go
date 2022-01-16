package forum

import "net/http"

func Server() {
	http.HandleFunc("/", HandleDefault)
	http.ListenAndServe(":8080", nil)	
}
