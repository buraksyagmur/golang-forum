package forum

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// fmt.Println("get home")
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/index.gohtml")
		// tpl, err := template.ParseFiles("./templates/index.gohtml")
		if err != nil {
			log.Fatal(err)
		}

		pos := displayPostsAndComments()

		err = tpl.ExecuteTemplate(w, "index.gohtml", pos)
		if err != nil {
			http.Error(w, "Executing Error", http.StatusInternalServerError)
		}
	}
	if r.Method == "POST" {
		newPost(r)
		addOne(r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/login.gohtml")
		if err != nil {
			log.Fatal(err)
		}

		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
	if r.Method == "POST" {
		processLogin(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/register.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		tpl.ExecuteTemplate(w, "register.gohtml", nil)
	}
	if r.Method == "POST" {
		regNewUser(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LogoutHanler(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		processLogout(w, r)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/notFound.gohtml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
// }
