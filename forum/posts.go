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
	// fmt.Println(postCon)

	stmt, err := db.Prepare("INSERT INTO post (postID, author, title, content, category, datetime, likes, dislikes) VALUES (?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(1, "DC", postTitle, postCon, "testCat", time.Now(), 7, 10)

	// test
	var pID int
	var author string
	var title string
	var content string
	var category string
	var datetime time.Time
	var likes int
	var dislikes int

	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&pID, &author, &title, &content, &category, &datetime, &likes, &dislikes)
		fmt.Printf("Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", pID, author, title, content, category, datetime, likes, dislikes)
	}
}

func displayPostsAndComments() {

}
