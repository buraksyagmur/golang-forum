package forum

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type comment struct {
	CommentID      int
	Content        string
	CommentTime    time.Time
	CommentTimeStr string
	Likes          int
	Dislikes       int
}

type post struct {
	PostID      int
	username    string
	Title       string
	Content     string
	Category    string
	PostTime    time.Time
	PostTimeStr string
	Likes       int
	Dislikes    int
	Comments    []comment
}

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
	idNumOfLikesStr := r.PostForm.Get("like")
	idNumOfDislikesStr := r.PostForm.Get("dislike")

	if idNumOfLikesStr != "" {
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
		stmt.Exec("DC", postTitle, postCon, postCatStr, time.Now(), 0, 0)
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

}

func displayComments(postID int) []comment {
	var coms []comment
	rows, err := db.Query(`
	SELECT * 
	FROM posts
	LEFT JOIN comments
		ON posts.postID = comments.postID
	WHERE comments.postID = ?
	`, postID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var com comment
		rows.Scan(&(com.CommentID), &(com.Content), &(com.CommentTime), &(com.Likes), &(com.Dislikes))
		com.CommentTimeStr = com.CommentTime.Format("Mon 02-01-2006 15:04:05")
		coms = append(coms, com)
	}

	return coms
}

func displayPostsAndComments() []post {
	// if filtered

	var pos []post
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var po post
		rows.Scan(&(po.PostID), &(po.username), &(po.Title), &(po.Content), &(po.Category), &(po.PostTime), &(po.Likes), &(po.Dislikes))
		po.PostTimeStr = po.PostTime.Format("Mon 02-01-2006 15:04:05")
		fmt.Printf("Display Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", po.PostID, po.username, po.Title, po.Content, po.Category, po.PostTimeStr, po.Likes, po.Dislikes)

		po.Comments = displayComments(po.PostID)
		pos = append(pos, po)
	}
	// for p := 0; p < len(pos); p++ {
	// }

	return pos
}

// func filterPost(r *http.Request) {
// 	r.ParseForm()
// }
