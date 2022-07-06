package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}
type Block struct {
	Pos  int
	Data BookCheckout
	Timestamp
	Hash
	PrevHash
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", getBlockchain).Methods("GET")
	// r.Handle("/", writeBlock).Methods("POST")
	// r.Handle("/new", newBook).Methods("POST")
	log.Println("Listening on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
