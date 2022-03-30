package forum

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}
func LogoutHanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
