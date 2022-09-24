package main

import (
	"fmt"
	"os"

	"github.com/brunno98/PersonApi/app/controller"
	"github.com/brunno98/PersonApi/app/infraestructure"
	"github.com/brunno98/PersonApi/app/model"
	"github.com/brunno98/PersonApi/app/service"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO: implementar logs
		fmt.Println("erro ao se conectar com o banco")
		os.Exit(1)
	}
	db.AutoMigrate(&model.Person{}) // TODO: Encontrar um lugar melhor para deixar esse codigo

	// database := infraestructure.InMemoryDB{}
	database := infraestructure.PersonRespositoryPostgres{DB: db}
	service := service.PersonService{IPersonRepository: &database}
	controller := controller.PersonController{IPersonService: &service}
	router := gin.Default()

	router.GET("/person/:id", controller.GetPersonById)
	router.POST("/person", controller.SavePerson)
	router.PUT("/person/:id", controller.UpdatePerson)
	router.DELETE("/person/:id", controller.DeletePerson)

	router.Run(":8080")
}
