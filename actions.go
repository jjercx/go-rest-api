package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// MovieIndex handler
func MovieIndex(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
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
	json.NewEncoder(w).Encode(movies)
}

// MovieShow handler
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	fmt.Fprintf(w, "Mostrando la pelicula %s", movieID)
}
