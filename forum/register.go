package forum

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// type user struct {
// 	Username string
// 	email    string
// 	password []byte
// 	access   int
// 	loggedIn bool
// }

func regNewUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := []byte(r.PostForm.Get("password"))

	// check if exists
	
	hash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO users (username, email, password, access, loggedIn) VALUES (?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(username, email, hash, 0, true)

	// test
	var u string
	var e string
	var p []byte
	var a int
	var l bool

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u, &e, &p, &a, &l)

	}
	fmt.Printf("uname: %s e: %s pw: %s, ac: %d, log: %t\n", u, e, p, a, l)

	sid := uuid.NewV4()
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  sid.String(),
		MaxAge: 1800,
	})
	// fmt.Println(sid.String())

	stmt, err = db.Prepare("INSERT INTO sessions (sessionID, username) VALUES (?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(sid.String(), username)
}
