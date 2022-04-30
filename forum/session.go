package forum

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func loggedIn(r *http.Request) bool {
	_, err := r.Cookie("session")
	if err != nil {
		return false
	}
	// var sid string
	// rows, err := db.Query("SELECT username, sessionID FROM sessions WHERE sessionID = ?;", c.Value)
	// if err != nil {
	// 	http.Error(w, "Error when verifying logged in status", http.StatusInternalServerError)
	// 	return false
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&sid)
	// }
	return true
}

func processLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	uname := r.PostForm.Get("username")
	pw := r.PostForm.Get("password")
	fmt.Printf("login u: %s: , login pw: %s\n", uname, pw)

	var unameDB string
	var hashDB []byte

	fmt.Printf("%s trying to Login\n", uname)
	rows, err := db.Query("SELECT username, password FROM users WHERE username = ?;", uname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&unameDB, &hashDB)
	}

	// test hash
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 10)

	fmt.Printf("unameDB: %s , hashDB: %s\n", unameDB, hashDB)
	err = bcrypt.CompareHashAndPassword(hashDB, []byte(pw))
	// fmt.Printf("DB pw: %s, entered: %s\n", hashDB, pw)
	fmt.Printf("DB pw: %s, entered: %s\n", hashDB, hash)
	if err != nil {
		http.Error(w, "Username or Password not matched, please try again", http.StatusForbidden)
		http.Redirect(w, r, "/login", http.StatusSeeOther) // not working
		return
	}
	fmt.Printf("%s (name from DB) Login successfully\n", unameDB)

	sid := uuid.NewV4()
	fmt.Printf("login sid: %s\n", sid)
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  sid.String(),
		MaxAge: 1800,
	})

	forumUser.Username = unameDB
	forumUser.Access = 1
	forumUser.LoggedIn = true
	fmt.Printf("%s forum User Login\n", forumUser.Username)

	stmt, err := db.Prepare("UPDATE users SET loggedIn = ? WHERE username = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(true, forumUser.Username)
}

func processLogout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	stmt, err := db.Prepare("DELETE FROM sessions WHERE sessionID=?")
	if c != nil {
		fmt.Printf("cookie sid to be removed (have value): %s\n", c.Value)

		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec(c.Value)
	}

	// test
	var sessionID string
	rows, err := db.Query("SELECT * FROM sessions")
	for rows.Next() {
		rows.Scan(&sessionID)
	}
	fmt.Printf("cookie sid removed (should be empty): %s\n", sessionID) // empty is correct

	_, err = r.Cookie("session")
	if err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	}
	fmt.Printf("%s Logout\n", forumUser.Username)

	stmt, err = db.Prepare("UPDATE users SET loggedIn = ? WHERE username = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(false, forumUser.Username)
	// test
	fmt.Printf("forumUser username: %s\n", forumUser.Username)
	fmt.Printf("Access should be 1: %d\n", forumUser.Access)

	forumUser = user{}
	fmt.Printf("forumUser username should be empty: %s\n", forumUser.Username)
	fmt.Printf("forumUser Access should be 0 (nil value): %d\n", forumUser.Access)
}
