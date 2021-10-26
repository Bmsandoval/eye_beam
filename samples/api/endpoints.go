package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)["id"]

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var person Person
	if val, ok := people[id]; ok {
		w.WriteHeader(http.StatusFound)
		person = val
	} else {
		w.WriteHeader(http.StatusNotFound)
		person = Person{}
	}

	json.NewEncoder(w).Encode(person)
	return
}

func ListPersonsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}