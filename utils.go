package main

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var validInterPage = regexp.MustCompile(`\[([A-Za-z]+)\]`)

func replacePageLink(body []byte) []byte {
	title := regexp.MustCompile(`(\[|\])`).ReplaceAll(body, []byte(""))
	link := `<a href="/view/` + string(title) + `">` + string(title) + `</a>`
	return []byte(link)
}

func interPageLink(p *Page) {
	newBody := template.HTML(validInterPage.ReplaceAllFunc([]byte(p.Body), replacePageLink))
	*p = Page{Title: p.Title, Body: newBody}

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	if r.URL.Path == "/" {
		return "", nil
	}
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}
