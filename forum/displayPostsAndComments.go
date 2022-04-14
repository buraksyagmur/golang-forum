package forum

import (
	"fmt"
	"log"
)

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
	fmt.Printf("forumUser username when display post: %s\n", forumUser.Username)
	var pos []post
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var po post
		rows.Scan(&(po.PostID), &(po.Author), &(po.Title), &(po.Content), &(po.Category), &(po.PostTime), &(po.Likes), &(po.Dislikes))
		po.PostTimeStr = po.PostTime.Format("Mon 02-01-2006 15:04:05")
		// fmt.Printf("Display Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", po.PostID, po.Author, po.Title, po.Content, po.Category, po.PostTimeStr, po.Likes, po.Dislikes)

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
