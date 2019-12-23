package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Struct(Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init Books var as slice book struct

var books []Book

//Get All books

func getBooks(w http.ResponseWriter, r *http.Request) {

}
func getBook(w http.ResponseWriter, r *http.Request) {

}
func createBook(w http.ResponseWriter, r *http.Request) {

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
}

func main() {
	//Init Router

	router := mux.NewRouter()

	//Mock data

	books = append(books, Book{ID: "1", Isbn: "1234", Title: "Book 1", Author: &Author{Firstname: "Tharindu", Lastname: "Yasantha"}})

	//Route handler

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/create", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
