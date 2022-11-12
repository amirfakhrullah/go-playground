package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

func postMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa((rand.Intn(1000000000)))
	movies = append(movies, newMovie)

	json.NewEncoder(w).Encode(newMovie)
}
