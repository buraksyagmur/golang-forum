package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func createSessionTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS session (sessionID INTEGER PRIMARY KEY,	username TEXT);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// test

}
func createUsersTable() {

}
func createPostsTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS post (postID INTEGER PRIMARY KEY AUTOINCREMENT, author VARCHAR(100), title VARCHAR(500), content VARCHAR(5000), category VARCHAR(50), datetime DATETIME, likes INTEGER, dislikes INTEGER);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()

	// test insert
	// stmt, err = db.Prepare("INSERT INTO post (postID, author, title, content, category, datetime, likes, dislikes) VALUES (?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// stmt.Exec(0, "user0", "test0", "testing0", "tech", time.Now(), 316, 777)

	// // test clear the test data before query
	// // stmt, err = db.Prepare("DELETE FROM post")
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // defer stmt.Close()
	// // stmt.Exec()

	// // test query
	// var pID int
	// var author string
	// var title string
	// var content string
	// var category string
	// var datetime time.Time
	// var likes int
	// var dislikes int

	// rows, err := db.Query("SELECT * FROM post")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	rows.Scan(&pID, &author, &title, &content, &category, &datetime, &likes, &dislikes)
	// 	fmt.Printf("Post: %d, by %s, title: %s, content: %s, in %s, at %v, with %d likes, and %d dislikes\n", pID, author, title, content, category, datetime, likes, dislikes)
	// }

	// clear the test data
	// stmt, err = db.Prepare("DELETE FROM post")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// stmt.Exec()
}

func InitDB() {
	db, _ = sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	createSessionTable()
	createUsersTable()
	createPostsTable()
}
