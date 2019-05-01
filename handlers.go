package main

import (
	"fmt"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	written, err := fmt.Fprint(w, "Main handler")
	if err != nil {
		log.Printf("%d bytes were written with error: %s", written, err)
	}
}

func handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainHandler)

	return mux
}
