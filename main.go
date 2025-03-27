package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html", "./tmpl/home.html"))

func main() {
	http.HandleFunc("/view/FrontPage", makeHandler(homeHandler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", makeHandler(homeHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
