package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.ListenAndServe(":8080", NewRouter())
	log.Fatal(server)
}
