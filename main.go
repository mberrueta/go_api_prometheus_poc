package main

import (
	"log"
	"net/http"
	"os"

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

var (
	bindAddress     = os.Getenv("BIND_ADDRESS")
	httpPort        = os.Getenv("HTTP_PORT")
	tlsDirectory    = os.Getenv("TLS_DIRECTORY")
	postgresHost    = os.Getenv("POSTGRES_HOST")
	postgresPort    = os.Getenv("POSTGRES_PORT")
	postgresDB      = os.Getenv("POSTGRES_DB")
	postgresUser    = os.Getenv("POSTGRES_USER")
	postgresPass    = os.Getenv("POSTGRES_PASS")
	postgresSSLMode = os.Getenv("POSTGRES_SSL_MODE")
)

func main() {
	if bindAddress == "" {
		bindAddress = ":50060"
	}
	if httpPort == "" {
		httpPort = ":8080"
	}
	// if postgresHost == "" {
	// 	log.Fatal("no postgres host specified, set POSTGRES_HOST env")
	// }
	// Init the mux router
	r := mux.NewRouter()

	// route handlers
	r.HandleFunc("/", someFunc)
	r.HandleFunc("/api/books", controllers.BooksIndex).Methods("GET")
	r.HandleFunc("/api/books", controllers.BooksCreate).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.BooksUpdate).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.BooksShow).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.BooksDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(httpPort, r))
}
