package controller

import (
	"ginpet/src/model"
	"ginpet/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePet(c *gin.Context) {
	var pet model.Pet

	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdPet, err := service.CREATE_PET(pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdPet, "message": "successfully created pet"})
}
