package main

import (
	"log"
	"net/http"

	"github.com/mberrueta/go_api_prometheus_poc/controllers"

	"github.com/gorilla/mux"
)

// ServerMux = HTTP request router = multiplexor = mux

// func main() {
// 	// http.HandleFunc("/", someFunc)
// 	// http.ListenAndServe(":8080", nil)
// 	http.ListenAndServe(":8080", mux())
// }

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!")) // slice of bytes
}

// func mux() *http.ServeMux {
// 	myMux := http.NewServeMux()
// 	// myMux.HandleFunc("/", someFunc)
// 	myMux.Handle("/", new(Handler))
// 	return myMux
// }

// type Handler struct {
// }

// func (this *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path
// 	log.Println(path)

// 	if path == "/not-exists" {
// 		w.Write([]byte("404: " + http.StatusText(404)))
// 	} else {
// 		w.Write([]byte("Hello!"))
// 	}
// }

func main() {
	// Init the mux router
	r := mux.NewRouter()

	// route handlers
	r.HandleFunc("/", someFunc)
	r.HandleFunc("/api/books", controllers.BooksIndex).Methods("GET")
	r.HandleFunc("/api/books", controllers.BooksCreate).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.BooksUpdate).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.BooksShow).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.BooksDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
