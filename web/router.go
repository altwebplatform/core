package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"log"
)

var templates = template.Must(template.ParseGlob("web/templates/*"))

func renderTemplate(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	name := params.ByName("template")
	if  len(name) == 0 {
		name = "main"
	}
	err := templates.ExecuteTemplate(w, name, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notFound(w http.ResponseWriter, req *http.Request) {
	log.Println("WEB - Not found: " + req.URL.Path)
}

func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", renderTemplate)
	router.GET("/dashboard/:template", renderTemplate)
	router.Handler("GET", "/static/*filepath",
		http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.NotFound = http.HandlerFunc(notFound)
	return router
}
