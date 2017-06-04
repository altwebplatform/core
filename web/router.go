package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"log"
	"github.com/altwebplatform/core/storage"
	"encoding/json"
	"io/ioutil"
	"strconv"
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
	http.NotFound(w, req)
}

func listModel(obj interface{}, key string, w http.ResponseWriter) {
	db := storage.SharedDB()
	model := db.Model(obj)
	rows, err := model.Limit(10).Rows()
	if err != nil {
		errorResponse(w, err)
		return
	}
	defer rows.Close()
	var resp []interface{}
	for rows.Next() {
		db.ScanRows(rows, &obj)
		resp = append(resp, obj)
	}
	if err := json.NewEncoder(w).Encode(map[string][]interface{}{key: resp}); err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getModel(obj interface{}, id uint64, w http.ResponseWriter) {
	db := storage.SharedDB()
	model := db.Model(obj)
	db = model.Find(obj,"id = ?", id)
	if db.Error != nil {
		errorResponse(w, db.Error)
		return
	}
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateModel(obj interface{}, id uint64, w http.ResponseWriter, req *http.Request) {

	// db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

	db := storage.SharedDB()
	model := db.Model(obj)
	db = model.Find(obj,"id = ?", id)
	if db.Error != nil {
		errorResponse(w, db.Error)
		return
	}
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createModel(obj interface{}, w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorResponse(w, err)
		return
	}
	if err := json.Unmarshal(body, obj); err != nil {
		errorResponse(w, err)
		return
	}
	db := storage.SharedDB()
	model := db.Model(obj)
	create := model.Create(obj)
	if err := create.Error; err != nil {
		errorResponse(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(map[string]string{"success": "true"}); err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteModel(obj interface{}, id uint64, w http.ResponseWriter, req *http.Request) {
	db := storage.SharedDB()
	if err := db.Where("id = ?", id).Delete(obj).Error; err != nil {
		errorResponse(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(map[string]string{"success": "true"}); err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	response := map[string]string{"success": "false", "message": err.Error()}
	if b, encodeErr := json.Marshal(response); encodeErr != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(b)
	}
	w.Header().Add("x-error", err.Error())
}

func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", renderTemplate)
	router.GET("/dashboard/:template", renderTemplate)


	router.POST("/api/v1/services", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		createModel(&storage.Service{}, w, req)
	})

	router.GET("/api/v1/services", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		listModel(&storage.Service{}, "services", w)
	})

	router.GET("/api/v1/services/:id", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id, err := strconv.ParseUint(params.ByName("id"), 10, 8)
		if err != nil {
			errorResponse(w, err)
			return
		}
		getModel(&storage.Service{}, id, w)
	})

	router.PUT("/api/v1/services/:id", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id, err := strconv.ParseUint(params.ByName("id"), 10, 8)
		if err != nil {
			errorResponse(w, err)
			return
		}
		updateModel(&storage.Service{}, id, w, req)
	})

	router.DELETE("/api/v1/services/:id", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id, err := strconv.ParseUint(params.ByName("id"), 10, 8)
		if err != nil {
			errorResponse(w, err)
			return
		}
		deleteModel(&storage.Service{}, id, w, req)
	})

	router.Handler("GET", "/static/*filepath",
		http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.NotFound = http.HandlerFunc(notFound)
	return router
}

