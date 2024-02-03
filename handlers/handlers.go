package handlers

import "net/http"

func Pippo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Hello Pippo!"))
}
