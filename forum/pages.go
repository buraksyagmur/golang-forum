package forum

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/index.gohtml")
	// tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "Executing Error", http.StatusInternalServerError)
	}
}
