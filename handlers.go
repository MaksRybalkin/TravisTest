package main

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

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

func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Release: %s\nCommit: %s\nRepo: %s\nDate: %s", RELEASE, COMMIT, REPO, DATE)
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	written, err := fmt.Fprintf(w, r.URL.String())
	if err != nil {
		log.Printf("%d bytes were written with error: %s", written, err)
	}
}

func handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/health", HealthHandler)

	return mux
}
