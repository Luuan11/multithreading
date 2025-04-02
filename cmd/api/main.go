package main

import (
	"gomultithreading/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Search API! Use /{cep} to search for a CEP."))
	})

	r.HandleFunc("/{cep}", handler.SearchCepHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
