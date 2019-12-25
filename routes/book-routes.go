package routes



import (
	"github.com/gorilla/mux"
	"github.com/restapi/controllers"
)

var BasicCrudRoutes = func (router *mux.Router ){

	router.HandleFunc("/api/books", controllers.getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.getBook).Methods("GET")
	router.HandleFunc("/api/books/create", controllers.createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.deleteBook).Methods("DELETE")
}