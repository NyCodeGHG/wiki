package handlers

import (
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"wiki/page"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	var pages []*page.Page
	dir, _ := ioutil.ReadDir("data")

	// Collect pages
	for _, file := range dir {
		fileName := file.Name()
		if !file.IsDir() && strings.HasSuffix(fileName, ".txt") {
			loadedPage, err := page.LoadPage(fileName[:len(fileName)-4])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			pages = append(pages, loadedPage)
		}
	}
	// Sort pages alphabetically
	sort.Slice(pages, func(i, j int) bool {
		switch strings.Compare(pages[i].Title, pages[j].Title) {
		case -1:
			return true
		case 1:
			return false
		}
		return pages[i].Title > pages[j].Title
	})

	err := templates.ExecuteTemplate(w, "list.gohtml", pages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
