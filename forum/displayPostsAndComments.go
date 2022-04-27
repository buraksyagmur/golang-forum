package forum

import (
	"log"
	"strings"
)

func displayComments(postID int) []comment {
	// fmt.Printf("postID: %d\n", postID)
	var coms []comment
	rows, err := db.Query("SELECT commentID, comments.[author] AS commentAuthor, comments.[postID], comments.[content], commentTime, comments.[likes], comments.[dislikes] FROM comments LEFT JOIN posts ON comments.postID = posts.postID	WHERE comments.postID = ?;", postID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var com comment
		rows.Scan(&(com.CommentID), &(com.Author), &(com.PostID), &(com.Content), &(com.CommentTime), &(com.Likes), &(com.Dislikes))
		com.CommentTimeStr = com.CommentTime.Format("Mon 02-01-2006 15:04:05")
		// fmt.Printf("CommentID: %d\n", com.CommentID)
		// fmt.Printf("Comment content: %s\n", com.Content)
		coms = append(coms, com)
	}

	return coms
}

func displayPostsAndComments() []post {
	// if filtered
	// fmt.Printf("forumUser username when display post: %s\n", forumUser.Username)
	var pos []post
	rows, err := db.Query("SELECT * FROM posts;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var po post
		rows.Scan(&(po.PostID), &(po.Author),&(po.Image), &(po.Title), &(po.Content), &(po.Category), &(po.PostTime), &(po.Likes), &(po.Dislikes))
		po.Category = strings.Trim(po.Category, "(")
		po.Category = strings.Trim(po.Category, ")")
		po.PostTimeStr = po.PostTime.Format("Mon 02-01-2006 15:04:05")
		// fmt.Printf("Display Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", po.PostID, po.Author, po.Title, po.Content, po.Category, po.PostTimeStr, po.Likes, po.Dislikes)

		po.Comments = displayComments(po.PostID)
		pos = append(pos, po)
	}
	return pos
}

// func filterPost(r *http.Request) {
// 	r.ParseForm()
// }
