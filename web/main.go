package web

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("web/templates/*"))

func renderTemplate(w http.ResponseWriter, req *http.Request) {
	name := "main"
	if len(req.URL.Path) > 1 {
		name = req.URL.Path[1:]
	}
	err := templates.ExecuteTemplate(w, name, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Start(connect string) {
	log.Println("Starting AWP server on " + connect + "...")
	staticServer := http.FileServer(http.Dir("static"))
	http.Handle("/css/", staticServer)
	http.Handle("/js/", staticServer)
	http.Handle("/vendor/", staticServer)
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(connect, nil)
	if err != nil {
		log.Fatal("Error starting AWP web dashboard: ", err)
	}
}
