/*
config: Contém funções relacionadas à configuração da aplicação, como a conexão com o banco de dados.

controller: Define controladores que manipulam as solicitações HTTP e interagem com os serviços.

helper: Fornece funções utilitárias, como a função ErrorPanic para lidar com erros e panics de maneira concisa.

model: Contém definições de estrutura para modelos de dados, como a estrutura Tags.

repository: Lida com a comunicação entre a aplicação e o banco de dados para operações relacionadas à entidade "tags".

router: Configura as rotas da aplicação usando o framework Gin.

service: Implementa lógica de negócios para operações relacionadas à entidade "tags".
*/

package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/router"
	"golang-crud-gin/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Servidor Inicializado")

	// Conexao com o Database
	db := config.DatabaseConnection()
	validate := validator.New()

	// Migracao automatica das tabelas
	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repository.NewTagsREpositoryImpl(db)

	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)

	// Criando rotas
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888", // porta
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
