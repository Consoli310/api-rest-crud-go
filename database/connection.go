package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// ConnectDatabase inicializa a conexão com o banco de dados
func ConnectDatabase() {
	var err error

	dsn := "root:1234567@tcp(localhost:3306)/product_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

// GetDatabase retorna a conexão com o banco de dados
func GetDatabase() *gorm.DB {
	return db
}
