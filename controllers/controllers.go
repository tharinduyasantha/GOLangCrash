package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/restapi/models"
	"github.com/restapi/utils"

	"github.com/gorilla/mux"
)

var books []models.Book

//Get All books
func getBooks(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book
func getBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	// Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}
func createBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	utils.ParseBody(router, book)

	_ = json.NewDecoder(router.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update books
func updateBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(router.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

//Delete book
func deleteBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
