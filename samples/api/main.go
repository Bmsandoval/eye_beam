package main

import (
	"github.com/gorilla/mux"
	"log"

	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", ListPersonsEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")

	svr := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			router.ServeHTTP(w, r)
		}),
	}

	log.Println("Starting Server...")
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatalf("error listening: %v", err.Error())
	}
}