package forum

import "net/http"

func Server() {
	http.ListenAndServe("8080", nil)	
}
