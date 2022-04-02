package forum

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	id       int
	username string
	email    string
	password []byte
	access   int
	loggedIn bool
}

func regNewUser(r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := []byte(r.FormValue("password"))
	// fmt.Printf("uname: %s, email: %s, pw: %s", newUser.username, newUser.email, newUser.password)
	hash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO users (id, username, email, password, access, loggedIn) VALUES (?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(1, username, email, hash, 0, true)

	// test
	var i int
	var u string
	var e string
	var p []byte
	var a int
	var l bool

	rows, err := db.Query("SELECT * FROM users")

	for rows.Next() {
		rows.Scan(&i, &u, &e, &p, &a, &l)
	}

	fmt.Printf("%d uname: %s e: %s pw: %s, ac: %d, log: %t", i, u, e, p, a, l)
}
