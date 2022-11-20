package routers

import (
	"github.com/amirfakhrullah/go-playground/tree/main/mongo-api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies", controllers.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}/mark", controllers.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", controllers.DeleteAllMovies).Methods("DELETE")

	return r
}