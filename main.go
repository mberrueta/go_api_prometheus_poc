package main

import (
	"log"
	"net/http"
)

// ServerMux = HTTP request router = multiplexor = mux

func main() {
	// http.HandleFunc("/", someFunc)
	// http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", mux())
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!")) // slice of bytes
}

func mux() *http.ServeMux {
	myMux := http.NewServeMux()
	// myMux.HandleFunc("/", someFunc)
	myMux.Handle("/", new(Handler))
	return myMux
}

type Handler struct {
}

func (this *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)

	if path == "/not-exists" {
		w.Write([]byte("404: " + http.StatusText(404)))
	} else {
		w.Write([]byte("Hello!"))
	}
}
