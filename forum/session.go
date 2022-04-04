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

func processLogin(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	uname := r.PostForm.Get("username")
	pw := r.PostForm.Get("password")
	fmt.Printf("%s: %s", uname, pw)

	// stmt, err := db.Prepare()
}

func LogoutHanler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	}
}
