package web

import (
	"log"
	"net/http"
)

func Start(connect string) {
	log.Println("Starting AWP server on " + connect + "...")
	log.Fatal(http.ListenAndServe(connect, CreateRouter()))
}
