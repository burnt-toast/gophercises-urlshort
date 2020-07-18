package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/burnt-toast/gophercises-urlshort"
)

func main() {
	mux := defaultMux()

	pathToUrls := map[string]string{
		"tech": "https://techcrunch.com/",
		"cats": "https://www.google.com/search?q=cat+memes",
	}

	mapHandler := urlshort.MapHandler(pathToUrls, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", googleForwarder)
	return mux
}

func googleForwarder(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://google.com", http.StatusFound)
}
