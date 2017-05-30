package main

import (
	"log"
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

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

func main() {
	connect := ":8080"
	log.Println("Starting server on " + connect + "...")
	staticServer := http.FileServer(http.Dir("static"))
	http.Handle("/css/", staticServer)
	http.Handle("/js/", staticServer)
	http.Handle("/vendor/", staticServer)
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(connect, nil)
	if err != nil {
		log.Println("Error starting: ", err)
		panic(err)
	}
}
