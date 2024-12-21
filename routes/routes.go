package routes

import (
	"api-crud-go/handlers"
	"github.com/gorilla/mux"
)

func ConfigureRoutes(router *mux.Router) {
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
}
