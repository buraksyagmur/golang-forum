package forum

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func processPostAndComment(r *http.Request) {
	// err := r.PostForm()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if filter
	// if post
	// if comment
}

func processPost(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	idNumOfLikesStr := r.PostForm.Get("po-like")
	idNumOfDislikesStr := r.PostForm.Get("po-dislike")

	if idNumOfLikesStr != "" {
		fmt.Printf("forumUser username when liking post: %s\n", forumUser.Username)
		idNumOfLikesStrSlice := strings.Split(idNumOfLikesStr, "-")
		updatePostID := idNumOfLikesStrSlice[0]
		numOfLikes, err := strconv.Atoi(idNumOfLikesStrSlice[1])
		if err != nil {
			log.Fatal(err)
		}
		numOfLikes++
		stmt, err := db.Prepare("UPDATE posts SET likes = ?	WHERE postID = ?;")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec(numOfLikes, updatePostID)
	} else if idNumOfDislikesStr != "" {
		fmt.Printf("forumUser username when disliking post: %s\n", forumUser.Username)
		idNumOfDislikesStrSlice := strings.Split(idNumOfDislikesStr, "-")
		updatePostID := idNumOfDislikesStrSlice[0]
		numOfDislikes, err := strconv.Atoi(idNumOfDislikesStrSlice[1])
		if err != nil {
			log.Fatal(err)
		}
		numOfDislikes++
		stmt, err := db.Prepare("UPDATE posts SET dislikes = ? WHERE postID = ?;")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec(numOfDislikes, updatePostID)
	} else {
		fmt.Printf("forumUser username when inserting new post: %s\n", forumUser.Username)
		postCon := r.PostForm.Get("postContent")
		postTitle := r.PostForm.Get("postTitle")
		postCat := r.PostForm["postCat"]
		// fmt.Println(postCon)

		stmt, err := db.Prepare("INSERT INTO posts (author, title, content, category, postTime, likes, dislikes) VALUES (?,?,?,?,?,?,?);")
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
		stmt.Exec(forumUser.Username, postTitle, postCon, postCatStr, time.Now(), 0, 0)
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
}

func processComment(r *http.Request) {
	r.ParseForm()
	idNumOfLikesStr := r.PostForm("com-like")
}
