package forum

import (
	"fmt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register")
}
