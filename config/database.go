package config

import (
	"fmt"
	"golang-crud-gin/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "" // Nome do host
	port     =  // Numero da porta
	user     = "" // Nome do usuario
	password = "" // Senha
	dbName   = "" // Nome do banco
)


// Cria conexao com o banco de dados
func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
