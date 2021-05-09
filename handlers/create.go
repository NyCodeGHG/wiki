package handlers

import (
	"errors"
	"net/http"
	"wiki/page"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodGet:
		err = templates.ExecuteTemplate(w, "create.gohtml", nil)
	case http.MethodPost:
		title := r.FormValue("entry_name")
		if existsPage(title) {
			http.Error(w, errors.New("entry already exists").Error(), http.StatusBadRequest)
			return
		}

		body := r.FormValue("body")
		newPage := &page.Page{
			Title: title,
			Body:  []byte(body),
		}
		err := newPage.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func existsPage(title string) bool {
	_, err := page.LoadPage(title)
	return err == nil
}
