package forum

import (
	"log"
)

func allForumUnames() []string {
	var allUnames []string
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var uname string
		rows.Scan(&uname)
		allUnames = append(allUnames, uname)
	}
	return allUnames
}

func filCatDisplayPostsAndComments(filCat string) []post {
	var pos []post
	// fmt.Printf("filCat is %s\n", filCat)
	filCat = "%(" + filCat + ")%"
	rows, err := db.Query("SELECT * FROM posts WHERE category LIKE ?;", filCat)
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
	return pos
}

func filAuthorDisplayPostsAndComments(filAuthor string) []post {
	var pos []post
	// fmt.Printf("filAuthor is %s\n", filAuthor)
	rows, err := db.Query("SELECT * FROM posts WHERE author LIKE ?;", filAuthor)
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
	return pos
}

func filLikedDisplayPostsAndComments() []post {
	var pos []post
	rows, err := db.Query("SELECT * FROM posts WHERE likes > 0;")
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
	return pos
}
