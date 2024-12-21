package handlers

import (
	"api-crud-go/database"
	"api-crud-go/models"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// GetProductByID busca um produto pelo ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Obtém o ID da URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db := database.GetDatabase() // Aqui usamos GetDatabase() para acessar a instância do banco
	var product models.Product

	// Consulta o produto no banco de dados
	result := db.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Retorna o produto encontrado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProducts retorna todos os produtos
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Conexão com o banco de dados
	db := database.GetDatabase() // Aqui usamos GetDatabase() para acessar a instância do banco
	var products []models.Product

	// Consulta todos os produtos
	result := db.Find(&products)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna os produtos encontrados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CreateProduct cria um novo produto
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db := database.GetDatabase()
	if err := db.Create(&p).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna status de sucesso
	w.WriteHeader(http.StatusCreated)
}

// UpdateProduct atualiza um produto existente
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db := database.GetDatabase()
	if err := db.Model(&models.Product{}).Where("id = ?", id).Updates(models.Product{Name: p.Name, Price: p.Price}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteProduct exclui um produto
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Conexão com o banco de dados
	db := database.GetDatabase()
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
