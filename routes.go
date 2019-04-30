package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{"root", "GET", "/", Index},
	Route{"movies", "GET", "/movies", MovieIndex},
	Route{"movie", "GET", "/movies/{id}", MovieShow},
	Route{"createMovie", "POST", "/movies", MovieCreate},
}

// Route represents http routes of the app
type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Routes is a list of Route structs
type Routes []Route

// NewRouter bootstraps the app HTTP routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, r := range routes {
		router.
			Name(r.Name).
			Methods(r.Method).
			Path(r.Path).
			Handler(r.Handler)
	}

	return router
}
