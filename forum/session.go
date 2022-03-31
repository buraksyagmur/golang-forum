package forum

import (
	"fmt"
	"log"
	"net/http"
)

func loggedIn(r *http.Request) bool {
	_, err := r.Cookie("session")
	if err != nil {
		return false
	}
	return true
}

func processLoginForm(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	uname := r.PostForm.Get("username")
	pw := r.PostForm.Get("password")
	fmt.Printf("%s: %s", uname, pw)
}

func LogoutHanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
