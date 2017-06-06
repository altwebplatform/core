package web

import (
	"encoding/json"
	"github.com/altwebplatform/core/storage"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	"github.com/jinzhu/gorm"
)

var templates = template.Must(template.ParseGlob("web/templates/*"))

type TypeCreator func() interface{}
type RowMapper func(*gorm.DB, *sql.Rows) interface{}

func renderTemplate(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	name := params.ByName("template")
	if len(name) == 0 {
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

func listModel(typeCreator TypeCreator, rowMapper RowMapper, key string, w http.ResponseWriter) {
	db := storage.SharedDB().Model(typeCreator())
	rows, err := db.Limit(10).Rows()
	if err != nil {
		errorResponse(w, err)
		return
	}
	defer rows.Close()
	var resp []interface{}

	//columns, err := rows.Columns()
	//if err != nil {
	//	errorResponse(w, err)
	//	return
	//}

	for rows.Next() {
		obj := rowMapper(db, rows)
		if err != nil {
			errorResponse(w, err)
			return
		}
		resp = append(resp, obj)
	}
	if err := json.NewEncoder(w).Encode(map[string][]interface{}{key: resp}); err != nil {
		errorResponse(w, err)
		return
	}
}

func getModel(typeCreator TypeCreator, w http.ResponseWriter, params httprouter.Params) {
	id, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		errorResponse(w, err)
		return
	}

	obj := typeCreator()
	db := storage.SharedDB()
	model := db.Model(obj)
	db = model.Find(obj, "id = ?", id)
	if db.Error != nil {
		errorResponse(w, db.Error)
		return
	}
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		errorResponse(w, err)
		return
	}
}

func updateModel(typeCreator TypeCreator, w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

	id, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		errorResponse(w, err)
		return
	}

	obj := typeCreator()
	db := storage.SharedDB()
	model := db.Model(obj)
	db = model.Find(obj, "id = ?", id)
	if db.Error != nil {
		errorResponse(w, db.Error)
		return
	}
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		errorResponse(w, err)
		return
	}
}

func createModel(typeCreator TypeCreator, w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorResponse(w, err)
		return
	}
	obj := storage.Service{}
	if err := json.Unmarshal(body, &obj); err != nil {
		errorResponse(w, err)
		return
	}
	db := storage.SharedDB()
	model := db.Model(&obj)
	create := model.Create(&obj)
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

func deleteModel(typeCreator TypeCreator, w http.ResponseWriter, params httprouter.Params) {
	//id, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	//if err != nil {
	//	errorResponse(w, err)
	//	return
	//}

	db := storage.SharedDB()
	if err := db.Delete(typeCreator(), "id = ?", params.ByName("id")).Error; err != nil {
		errorResponse(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(map[string]string{"success": "true"}); err != nil {
		errorResponse(w, err)
		return
	}
}

func errorResponse(w http.ResponseWriter, err error) {
	w.Header().Add("x-error", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	response := map[string]string{"success": "false", "message": err.Error()}
	if b, encodeErr := json.Marshal(response); encodeErr != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(b)
	}
}

func handleType(router *httprouter.Router, path string, createType TypeCreator) {
	router.POST("/api/v1/" + path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		createModel(createType, w, req)
	})
	router.GET("/api/v1/" + path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		listModel(createType, func(db *gorm.DB, rows *sql.Rows) interface{} {
			obj := &storage.Service{}
			db.ScanRows(rows, &obj)
			return obj
		},
			"services", w)
	})
	router.GET("/api/v1/" + path + "/:id", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		getModel(createType, w, params)
	})
	router.PUT("/api/v1/" + path + "/:id", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		updateModel(createType, w, req, params)
	})
	router.DELETE("/api/v1/" + path + "/:id", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		deleteModel(createType, w, params)
	})
}

func CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", renderTemplate)
	router.GET("/dashboard/:template", renderTemplate)

	handleType(router, "services", func() interface{} { return storage.Service{} })

	router.Handler("GET", "/static/*filepath",
		http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	router.NotFound = http.HandlerFunc(notFound)
	return router
}
