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

}
func createUsersTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT,username VARCHAR(300),email VARCHAR(500),password VARCHAR(500),access INTEGER,loggedIn BOOLEAN);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}
func createPostsTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts (postID INTEGER PRIMARY KEY AUTOINCREMENT, author VARCHAR(100), title VARCHAR(500), content VARCHAR(5000), category VARCHAR(50), datetime DATETIME, likes INTEGER, dislikes INTEGER);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func createCommentsTable() {
	// stmt, err := db.Prepare()
}

func InitDB() {
	db, _ = sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	createSessionTable()
	createUsersTable()
	createPostsTable()
	createCommentsTable()
}
