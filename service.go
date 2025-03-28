package main

import (
	"html/template"
	"log"
	"os"
	"sort"
	"strings"
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

	return &Page{Title: title, Body: template.HTML(bytes)}, nil
}

func getTopTenOfPage() ([]Page, error) {
	files, err := getFiles()

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})

	pages := selectFiles(files)

	return pages, nil
}

func getFiles() ([]os.FileInfo, error) {

	file, err := os.Open("./data/")

	if err != nil {
		return nil, err
	}

	files, err := file.Readdir(0)
	return files, err
}

func selectFiles(files []os.FileInfo) []Page {
	var pages []Page
	stopCount := 10
	if lenInfos := len(files); lenInfos < 10 {
		stopCount = lenInfos
	}
	for index := 0; index < stopCount; {
		name := files[index].Name()
		file, err := os.ReadFile("./data/" + name)
		if err != nil {
			log.Fatal(err)
			continue
		}

		page := &Page{Title: strings.Split(name, ".")[0], Body: template.HTML(file)}
		pages = append(pages, *page)
		index++
	}
	return pages
}
