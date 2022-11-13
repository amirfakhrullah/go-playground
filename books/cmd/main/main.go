package main

import (
	"log"
	"net/http"
	"github.com/amirfakhrullah/go-playground/tree/main/books/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
