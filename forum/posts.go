package forum

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func newPost(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	postCon := r.PostForm.Get("postContent")
	postTitle := r.PostForm.Get("postTitle")
	postCat := r.PostForm["postCat"]
	// fmt.Println(postCon)

	stmt, err := db.Prepare("INSERT INTO posts (username, title, content, category, postTime, likes, dislikes) VALUES (?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	postCatStr := ""
	for i := 0; i < len(postCat); i++ {
		if i != len(postCat)-1 {
			postCatStr += postCat[i] + "+"
		}
		if i == len(postCat)-1 {
			postCatStr += postCat[i]
		}
	}
	stmt.Exec("DC", postTitle, postCon, postCatStr, time.Now(), 7, 10)
	stmt.Exec("ST", postTitle, postCon, postCatStr, time.Now().Add(time.Minute*20), 3, 16)
	// test
	var pID int
	var username string
	var title string
	var content string
	var category string
	var postTime time.Time
	var likes int
	var dislikes int

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&pID, &username, &title, &content, &category, &postTime, &likes, &dislikes)
		fmt.Printf("Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", pID, username, title, content, category, postTime, likes, dislikes)
	}
}

func displayPostsAndComments() {

}
