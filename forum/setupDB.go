package forum

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func createSessionTable(db *sql.DB) {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS session (sessionID INTEGER PRIMARY KEY,	username TEXT);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// test

}
func createUsersTable(db *sql.DB) {

}
func createPostsTable(db *sql.DB) {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS post (postID INTEGER PRIMARY KEY, title TEXT, content TEXT,	datetime DATETIME);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()

	// test insert
	stmt, err = db.Prepare("INSERT INTO post (postID, title, content, datetime) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(0, "test0", "testing0", 0)

	// test query
	var pID int
	var title string
	var content string
	var datetime time.Time

	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&pID, &title, &content, &datetime)
		fmt.Printf("Post: %d, title: %s, content: %s, at %v\n", pID, title, content, datetime)
	}

}
func InitDB() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	createSessionTable(db)
	createUsersTable(db)
	createPostsTable(db)
}
