package main

func initialiseMovies() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1223",
		Title: "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "2324",
		Title: "The Black Panther",
		Director: &Director{
			Firstname: "Ryan",
			Lastname:  "Coogler",
		},
	})
}