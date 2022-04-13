package forum

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// var gh int
// type indexData struct {
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// gh++
		// fmt.Println(gh)
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/index.gohtml")
		// tpl, err := template.ParseFiles("./templates/index.gohtml")
		if err != nil {
			log.Fatal(err)
		}

		pos := displayPostsAndComments()
		allForumUnames := allForumUnames()

		data := struct {
			CurUser     user
			Posts       []post
			ForumUnames []string
		}{
			forumUser,
			pos,
			allForumUnames,
		}
		err = tpl.ExecuteTemplate(w, "index.gohtml", data)
		if err != nil {
			http.Error(w, "Executing Error", http.StatusInternalServerError)
		}
	}
	if r.Method == "POST" {
		r.ParseForm()
		filterCategory := r.PostForm.Get("category-filter")
		fmt.Printf("Filter Category %s\n", filterCategory)
		if filterCategory != "" {
			fmt.Printf("Filter Category %s\n", filterCategory)
		} else {

		}
		processPost(r)
		// processPostAndComment(r)
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

// func DeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	// for testing purpose
// 	if r.Method == "GET" {
// 		tpl, err := template.ParseFiles("./templates/delete.gohtml", "./templates/footer.gohtml", "./templates/header.gohtml")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tpl.ExecuteTemplate(w, "delete.gohtml", nil)
// 	}
// 	if r.Method == "POST" {
// 		deleteUser(r)
// 	}
// }

// func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/notFound.gohtml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
// }
