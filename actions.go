package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{
		"El Señor de los Anillos",
		2001,
		"Peter Jackson",
	},
	Movie{
		"Titanic",
		1998,
		"James Cameron",
	},
}

// Index just prints Hello World!
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// MovieIndex list all movies
func MovieIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

// MovieShow list details from an specific movie
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	fmt.Fprintf(w, "Mostrando la pelicula %s", movieID)
}

// MovieCreate adds a movie to the list
func MovieCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	newMovie := Movie{}
	err := decoder.Decode(&newMovie)
	if err != nil {
		w.WriteHeader(500)
	}
	defer r.Body.Close()

	movieCollection.Insert(newMovie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
