package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	fmt.Println("Starting server at port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
