package forum

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type post struct {
	postID      int
	username    string
	Title       string
	Content     string
	Category    string
	PostTime    time.Time
	PostTimeStr string
	Likes       int
	Dislikes    int
}

func newPost(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	postCon := r.PostForm.Get("postContent")
	postTitle := r.PostForm.Get("postTitle")
	postCat := r.PostForm["postCat"]
	// fmt.Println(postCon)
	// liked := r.PostForm.Get("like")
	// // disliked := r.PostForm.Get("dislike")
	// fmt.Println("-----------")
	// fmt.Println(liked)

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
	// stmt.Exec("ST", postTitle, postCon, postCatStr, time.Now().Add(time.Minute*20), 3, 16)

	// test
	// var pid int
	// var un string
	// var t string
	// var con string
	// var cat string
	// var pT time.Time
	// var l int
	// var d int

	// rows, err := db.Query("SELECT * FROM posts")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	rows.Scan(&pid, &un, &t, &con, &cat, &pT, &l, &d)
	// 	fmt.Printf("Post: %d, by %s, Title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", pid, un, t, con, cat, pT, l, d)
	// }
}

func addOne(r *http.Request) {
	numOfLikesStr := r.PostForm.Get("like")
	numOfLikes, err := strconv.Atoi(numOfLikesStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----------")
	fmt.Printf("likes: %d\n", numOfLikes)
	fmt.Println("-----------")

	// disliked := r.PostForm.Get("dislike")
}

func displayPostsAndComments() []post {
	var pos []post
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var po post
		rows.Scan(&(po.postID), &(po.username), &(po.Title), &(po.Content), &(po.Category), &(po.PostTime), &(po.Likes), &(po.Dislikes))
		po.PostTimeStr = po.PostTime.Format("Mon 02-01-2006 15:04:05")
		fmt.Printf("Display Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", po.postID, po.username, po.Title, po.Content, po.Category, po.PostTimeStr, po.Likes, po.Dislikes)
		pos = append(pos, po)
	}
	return pos
}
