package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func createUsersTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (username VARCHAR(50) PRIMARY KEY, email VARCHAR(50), password VARCHAR(100), access INTEGER, loggedIn BOOLEAN);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func createSessionsTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS sessions (sessionID VARCHAR(50) PRIMARY KEY, username VARCHAR(50), FOREIGN KEY(username) REFERENCES users(username));")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func createPostsTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts (postID INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(50), title VARCHAR(300), content VARCHAR(2000), category VARCHAR(50), postTime DATETIME, likes INTEGER, dislikes INTEGER, FOREIGN KEY(username) REFERENCES users(username));")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func createCommentsTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS comments (commentID INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(50), postID INTEGER, content VARCHAR(2000), commentTime DATETIME, likes INTEGER, dislikes INTEGER, FOREIGN KEY(username) REFERENCES users(username), FOREIGN KEY(postID) REFERENCES posts(postID));")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func InitDB() {
	db, _ = sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	createSessionsTable()
	createUsersTable()
	createPostsTable()
	createCommentsTable()
}
