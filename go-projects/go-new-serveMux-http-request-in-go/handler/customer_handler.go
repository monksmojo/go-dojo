package handler

import "net/http"

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userID")
	w.Write([]byte("userId is: " + userId))
}
	