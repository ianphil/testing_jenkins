package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// test out of sample
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
