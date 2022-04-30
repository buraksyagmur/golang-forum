package forum

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type mainPageData struct {
	Userinfo    user
	Posts       []post
	ForumUnames []string
}

var urlPost string

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	curUser := checkCookie(r)

	// // test
	// var whichUser string
	// var logInOrNot bool
	// rows, err := db.Query("SELECT username, loggedIn FROM users WHERE username = ?;", curUser.Username)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&whichUser, &logInOrNot)
	// }
	// fmt.Printf("HomeHandler:: login user: %s, login status: %v\n", whichUser, logInOrNot)

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/header2.gohtml", "./templates/footer.gohtml", "./templates/index.gohtml")
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
			Userinfo:    curUser,
			ForumUnames: allForumUnames,
		}
		// fmt.Println("---------", forumUser)
		err = tpl.ExecuteTemplate(w, "index.gohtml", data)
		if err != nil {
			http.Error(w, "Executing Error", http.StatusInternalServerError)
		}
	}
	if r.Method == http.MethodPost {
		processPost(r, curUser)
		processComment(r, curUser)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logged in", loggedIn(r))
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
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	//	if loggedIn(r) {
	processLogout(w, r)
	//	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	curUser := checkCookie(r)
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/header2.gohtml", "./templates/post.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		strID := r.FormValue("postdetails")
		PostIdFromHTML, err := strconv.Atoi(strID)
		if err != nil {
			os.Exit(0)
		}
		// fmt.Println(PostIdFromHTML, "---------")
		pos := displayPostsAndComments()

		allForumUnames := allForumUnames()
		var Chosen []post
		for i := 0; i < len(pos); i++ {
			if pos[i].PostID == PostIdFromHTML {
				Chosen = append(Chosen, pos[i])
			}
		}

		urlPost = "postpage?postdetails=" + strID + "&postdetails=" + Chosen[0].Title
		data := mainPageData{
			Posts:       Chosen,
			Userinfo:    curUser,
			ForumUnames: allForumUnames,
		}

		err = tpl.ExecuteTemplate(w, "post.gohtml", data)
		if err != nil {
			http.Error(w, "Executing Error", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {

		processPost(r, curUser)
		processComment(r, curUser)
		http.Redirect(w, r, urlPost, http.StatusSeeOther)
	}
}

func CategoryPageHandler(w http.ResponseWriter, r *http.Request) {
	curUser := checkCookie(r)
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/header2.gohtml", "./templates/footer.gohtml", "./templates/categories.gohtml")
		if err != nil {
			log.Fatal(err)
		}

		var pos []post
		category := r.FormValue("categoryAllPosts")
		pos = filCatDisplayPostsAndComments(category)

		allForumUnames := allForumUnames()
		data := mainPageData{
			Posts:       pos,
			Userinfo:    curUser,
			ForumUnames: allForumUnames,
		}
		// fmt.Println("---------", forumUser)
		err = tpl.ExecuteTemplate(w, "categories.gohtml", data)
		if err != nil {
			http.Error(w, "Executing Error", http.StatusInternalServerError)
		}
	}
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
