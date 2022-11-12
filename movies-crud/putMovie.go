package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func putMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			updatedMovie.ID = params["id"]
			movies = append(movies, updatedMovie)

			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
}
