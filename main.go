package main

import (
	"api-crud-go/database"
	"api-crud-go/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializa o banco de dados
	database.ConnectDatabase()

	// Configura o roteador
	router := mux.NewRouter()
	routes.ConfigureRoutes(router)

	// Inicia o servidor
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
