package main

import (
	"log"
	"net/http"
)

func main() {

	// match url and handler
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	// register the file server
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
