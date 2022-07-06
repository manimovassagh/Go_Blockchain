package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
}

type Book struct {
}

type BookCheckout struct {
}

type Blockchain struct {
	blocks []*Block
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.Handle("/", writeBlock).Methods("POST")
	r.Handle("/new", newBook).Methods("POST")
	log.Println("Listening on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
