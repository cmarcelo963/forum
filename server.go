package forum

import "net/http"


//This runs the server sends whatever information we want to show the user
func Server() {
	http.Handle("/index.css", http.FileServer(http.Dir("../static")))
	http.Handle("/index.js", http.FileServer(http.Dir("../static")))
	fileServer := http.FileServer(http.Dir("../static/img"))
	http.Handle("/img/", http.StripPrefix("/img", fileServer))
	http.HandleFunc("/", HandleDefault)
	http.HandleFunc("/sign-up", HandleSignUpRequest)
	http.HandleFunc("/login", HandleLoginRequest)
	http.HandleFunc("/new-post", HandleNewPostRequest)
	http.HandleFunc("/filter-posts", HandleFilterRequest)
	http.HandleFunc("/select-post", HandleSelectPost)
	http.HandleFunc("/new-comment", HandleNewCommentRequest)
	http.HandleFunc("/like", HandleNewLikeRequest)
	http.HandleFunc("/dislike", HandleNewLikeRequest)
	http.ListenAndServe(":8080", nil)
}
