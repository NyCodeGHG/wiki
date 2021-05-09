package handlers

import (
	"html/template"
	"net/http"
	"regexp"
	"wiki/page"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".gohtml", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
