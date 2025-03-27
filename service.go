package main

import (
	"html/template"
	"os"
)

type Page struct {
	Title string
	Body  template.HTML
}

func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	return os.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	body := template.HTML(bytes)

	return &Page{Title: title, Body: body}, nil
}
