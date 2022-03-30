package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createSessionTable(db *sql.DB) {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS session (sessionID INTEGER PRIMARY KEY,	username TEXT);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
}
func createUsersTable(db *sql.DB) {

}
func createPostsTable(db *sql.DB) {

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
