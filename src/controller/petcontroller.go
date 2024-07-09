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
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "successfully created pet", "data": createdPet})
}

func UpdatePet(c *gin.Context) {
	id := c.Param("id")
	var pet model.Pet

	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatePet, err := service.UPDATE_PET(id, pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully update pet", "data": updatePet})
}

func DeletePet(c *gin.Context) {
	id := c.Param("id")

	data, err := service.DELETE_PET(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully delete pet", "removed data": data})
}

func GetByIdPet(c *gin.Context) {
	id := c.Param("id")

	data, err := service.GET_BY_ID_PET(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successsfully fetch pet", "fetched data": data})
}
