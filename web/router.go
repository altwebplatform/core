package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"log"
	"github.com/altwebplatform/core/storage"
	"encoding/json"
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

func listServices(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	listModel(&storage.Service{}, "services", w)
}

func listModel(obj interface{}, key string, w http.ResponseWriter) {
	db := storage.SharedDB()
	model := db.Model(obj)
	rows, err := model.Limit(10).Rows()
	if err != nil {
		errorResponse(w, err)
	}
	defer rows.Close()
	var resp []interface{}
	for rows.Next() {
		db.ScanRows(rows, &obj)
		resp = append(resp, obj)
	}
	b, err := json.Marshal(map[string][]interface{}{key: resp})
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func errorResponse(writer http.ResponseWriter, err error) {
	writer.Write([]byte(err.Error()))
	writer.WriteHeader(http.StatusInternalServerError)
}

func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", renderTemplate)
	router.GET("/dashboard/:template", renderTemplate)

	router.GET("/api/v1/service/list", listServices)

	router.Handler("GET", "/static/*filepath",
		http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.NotFound = http.HandlerFunc(notFound)
	return router
}
