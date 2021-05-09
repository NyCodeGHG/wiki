package handlers

import (
	"net/http"
	"wiki/page"
)

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
