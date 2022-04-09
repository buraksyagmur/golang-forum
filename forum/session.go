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
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sid := uuid.NewV4()
	fmt.Println(sid)
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  sid.String(),
		MaxAge: 1800,
	})
}

func processLogout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	fmt.Printf("cookie sid to be removed: %s", c.Value)
	stmt, err := db.Prepare("DELETE FROM sessions WHERE sessionID=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(c.Value)

	//test
	var sessionID string
	rows, err := db.Query("SELECT * FROM sessions")
	for rows.Next() {
		rows.Scan(&sessionID)
	}
	fmt.Printf("cookie sid removed: %s", sessionID) // empty is correct

	_, err = r.Cookie("session")
	if err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	}

}
