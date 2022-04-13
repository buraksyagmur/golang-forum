package forum

import (
	"log"
	"net/http"
)

func deleteUser(r *http.Request) {
	r.ParseForm()
	dUser := r.PostForm.Get("delete")
	stmt, err := db.Prepare("DELETE FROM users WHERE username = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(dUser)
}
