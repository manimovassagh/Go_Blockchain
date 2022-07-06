package main

import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.Handle("/", writeBlock).Methods("POST")
	r.Handle("/new", newBook).Methods("POST")

}
