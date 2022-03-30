package forum

import (
	"fmt"
	"html/template"
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if loggedIn(r) {
			// go to other page
		}
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/login.gohtml")
		if err != nil {
			log.Fatal(err)
		}

		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
	if r.Method == "POST" {

	}

}
func LogoutHanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
