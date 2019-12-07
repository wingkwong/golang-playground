package main

import (
	"fmt"
	"net/http"

	urlshorter "github.com/golang-playground/url-shorter"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/google": "https://google.com",
		"/github": "https://github.com",
	}

	mapHandler := urlshorter.MapHandler(pathsToUrls, mux)

	http.ListenAndServe(":8888", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
