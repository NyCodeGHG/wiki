package main

import (
	"log"
	"net/http"
	"wiki/handlers"
)

func main() {
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))
	http.HandleFunc("/list/", handlers.ListHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/list/", http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
