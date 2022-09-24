package main

import (
	"github.com/brunno98/PersonApi/app/controller"
	"github.com/brunno98/PersonApi/app/infraestructure"
	"github.com/brunno98/PersonApi/app/service"
	"github.com/gin-gonic/gin"
)

func main() {
	database := infraestructure.InMemoryDB{}
	service := service.PersonService{IPersonRepository: &database}
	controller := controller.PersonController{IPersonService: &service}
	router := gin.Default()

	router.GET("/person/:id", controller.GetPersonById)
	router.POST("/person", controller.SavePerson)
	router.PUT("/person/:id", controller.UpdatePerson)
	router.DELETE("/person/:id", controller.DeletePerson)

	router.Run(":8080")
}
