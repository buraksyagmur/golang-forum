package forum

import (
	"fmt"
	"log"
	"net/http"
)

func loggedIn(r *http.Request) bool {
	_, err := r.Cookie("session")
	if err != nil {
		return false
	}
	return true
}

func processLogin(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	uname := r.PostForm.Get("username")
	pw := r.PostForm.Get("password")
	fmt.Printf("%s: %s", uname, pw)

	// stmt, err := db.Prepare()
}

func LogoutHanler(w http.ResponseWriter, r *http.Request) {
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
