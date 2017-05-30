package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	fmt.Println("STARTED!")
	err := http.ListenAndServe(":30080", HTTPProxy())
	if err != nil {
		log.Println("Error with HTTP Server: " + err.Error())
	}
	log.Println("Reloader server shutdown")
}

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func HTTPProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:80"
	}
	modifyResponse := func(resp *http.Response) error {
		resp.Header.Add("x-reloader", "true")
		return nil
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
}
