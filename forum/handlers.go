package forum

import (
	"html/template"
	"log"
	"net/http"
)

type mainPageData struct {
	Userinfo    user
	Posts       []post
	ForumUnames []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Yype", "text/html; charset=utf-8")
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/index.gohtml")
		// tpl, err := template.ParseFiles("./templates/index.gohtml")
		if err != nil {
			log.Fatal(err)
		}

		filCat := r.FormValue("category-filter")
		filAuthor := r.FormValue("author-filter")
		filLiked := r.FormValue("liked-post")

		var pos []post
		if filCat != "" {
			pos = filCatDisplayPostsAndComments(filCat)
		} else if filAuthor != "" {
			pos = filAuthorDisplayPostsAndComments(filAuthor)
		} else if filLiked != "" {
			pos = filLikedDisplayPostsAndComments()
		} else {
			pos = displayPostsAndComments()
		}

		allForumUnames := allForumUnames()

		data := mainPageData{
			Posts:       pos,
			Userinfo:    forumUser,
			ForumUnames: allForumUnames,
		}
		// fmt.Println("---------", forumUser)
		err = tpl.ExecuteTemplate(w, "header.gohtml", data)
		err = tpl.ExecuteTemplate(w, "index.gohtml", data)
		if err != nil {
			http.Error(w, "Executing Error", http.StatusInternalServerError)
		}
	}
	if r.Method == http.MethodPost {
		processPost(r)
		processComment(r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Yype", "text/html; charset=utf-8")
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/login.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
	if r.Method == http.MethodPost {
		processLogin(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Yype", "text/html; charset=utf-8")
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/register.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		tpl.ExecuteTemplate(w, "register.gohtml", nil)
	}
	if r.Method == http.MethodPost {
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
// 	if r.Method == http.MethodGet {
// 		tpl, err := template.ParseFiles("./templates/delete.gohtml", "./templates/footer.gohtml", "./templates/header.gohtml")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tpl.ExecuteTemplate(w, "delete.gohtml", nil)
// 	}
// 	if r.Method == http.MethodPost {
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
