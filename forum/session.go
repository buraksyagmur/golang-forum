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
	// get login data from form
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	uname := r.PostForm.Get("username")
	pw := r.PostForm.Get("password")
	fmt.Printf("login u: %s: , login pw: %s\n", uname, pw)

	// get user data from db
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

	// compare pw
	err = bcrypt.CompareHashAndPassword(hashDB, []byte(pw))
	// fmt.Printf("DB pw: %s, entered: %s\n", hashDB, pw)
	fmt.Printf("DB pw: %s, entered: %s\n", hashDB, hash)
	if err != nil {
		http.Error(w, "Username or Password not matched, please try again", http.StatusForbidden)
		http.Redirect(w, r, "/login", http.StatusSeeOther) // not working
		return
	}
	fmt.Printf("%s (name from DB) Login successfully\n", unameDB)

	// assign a cookie
	sid := uuid.NewV4()
	fmt.Printf("login sid: %s\n", sid)
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  sid.String(),
		MaxAge: 1800,
	})

	// forumUser.Username = unameDB
	// forumUser.Access = 1
	// forumUser.LoggedIn = true
	// fmt.Printf("%s forum User Login\n", forumUser.Username)

	// update the user's login status
	stmt, err := db.Prepare("UPDATE users SET loggedIn = ? WHERE username = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(true, unameDB)

	// insert a record into session table
	stmt, err = db.Prepare("INSERT INTO sessions (sessionID, username) VALUES (?, ?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(sid.String(), unameDB)

	//test
	// var whichUser string
	// var logInOrNot bool
	// rows, err = db.Query("SELECT username, loggedIn FROM users WHERE username = ?;", unameDB)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&whichUser, &logInOrNot)
	// }
	// fmt.Printf("login user: %s, login status: %v\n", whichUser, logInOrNot)
}

func processLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	var logoutUname string

	if err == nil {
		// get the username of the logout user
		rows, err := db.Query("SELECT username FROM sessions WHERE sessionID = ?;", c.Value)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&logoutUname)
		}
		fmt.Printf("Found user %s wants to logout", logoutUname)

		// delete sessionID from sessions db table
		stmt, err := db.Prepare("DELETE FROM sessions WHERE sessionID=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec(c.Value)
		fmt.Printf("cookie sid removed (have value): %s\n", c.Value)
	}

	// test
	var sessionID string
	rows, err := db.Query("SELECT * FROM sessions")
	for rows.Next() {
		rows.Scan(&sessionID)
	}
	fmt.Printf("cookie sid removed (should be empty): %s\n", sessionID) // empty is correct

	// delete browser's cookie
	_, err = r.Cookie("session")
	if err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	}
	fmt.Printf("%s Logout\n", logoutUname)

	stmt, err := db.Prepare("UPDATE users SET loggedIn = ? WHERE username = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(false, logoutUname)
}

func checkCookie(r *http.Request) user {
	var curUser user
	c, err := r.Cookie("session")
	// if there is a session cookie
	if err == nil {
		fmt.Printf("There is a cookie, sid: %s\n", c.Value)
		// get current username from session table
		var curUname string
		rows, err := db.Query("SELECT username FROM sessions WHERE sessionID = ?;", c.Value)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&curUname)
			// fmt.Printf("Found uname %s in sessions\n", curUname)
		}
		fmt.Printf("Found uname %s in sessions\n", curUname)
		rows, err = db.Query("SELECT username, image, email, access, loggedIN  FROM users WHERE username = ?;", curUname)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&curUser.Username, &curUser.Image, &curUser.Email, &curUser.Access, &curUser.LoggedIn)
			fmt.Printf("Found user %s in users, with login status %v\n", curUser.Username, curUser.LoggedIn)
		}
	}

	// test
	var whichUser string
	var logInOrNot bool
	rows, err := db.Query("SELECT username, loggedIn FROM users WHERE username = ?;", curUser.Username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&whichUser, &logInOrNot)
	}
	fmt.Printf("checkCookie:: login user: %s, login status: %v\n", whichUser, logInOrNot)

	return curUser
}
