package main

var movieCollection = getSession().DB("udemy-go-rest-api").C("movies")

// Movie represents a movie
type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

// Movies is a list of Movie structs
type Movies []Movie
