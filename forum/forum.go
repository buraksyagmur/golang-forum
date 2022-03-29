package forum

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/header.gohtml", "./template/footer.gohtml", "./template/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
