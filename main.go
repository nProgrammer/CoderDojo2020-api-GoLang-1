package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", block)
	r.HandleFunc("/hello", hello)

	http.ListenAndServe(":8000", r)
}

func block(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("PORT :8000 - blocked")
}

func hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello world!")
}
