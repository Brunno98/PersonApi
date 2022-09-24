package controller

import (
	"net/http"
	"strconv"

	"github.com/brunno98/PersonApi/app/domain"
	"github.com/brunno98/PersonApi/app/interfaces"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	interfaces.IPersonService
}

func (p *PersonController) GetPersonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	person, err := p.GetById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if person == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, person)
}

func (p *PersonController) SavePerson(c *gin.Context) {
	requestPerson := domain.Person{}
	c.ShouldBindJSON(&requestPerson)
	person, err := p.Save(&requestPerson)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}
	c.JSON(http.StatusOK, person)
}

func (p *PersonController) UpdatePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	person := domain.Person{}
	c.ShouldBindJSON(&person)

	if person.Id == 0 {
		person.Id = id
	}
	if person.Id != id {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"message": "id do path e id do body não correspondem."})
		return
	}

	if _, err := p.Update(&person); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (p *PersonController) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = p.Delete(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
