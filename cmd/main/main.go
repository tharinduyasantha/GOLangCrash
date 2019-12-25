package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restapi/routes"
)

func main() {
	
	router := mux.NewRouter()
	routes.BasicCrudRoutes(router)
	

	log.Fatal(http.ListenAndServe(":8000", router))
}
