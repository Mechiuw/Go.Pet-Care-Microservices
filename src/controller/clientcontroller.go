package controller

import (
	"net/http"

	"ginpet/src/model"
	"ginpet/src/service"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var client model.Client
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	createdClient, err := service.CREATE_CLIENT(client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdClient})
}

func update(c *gin.Context) {
	id := c.Param("id")
	var updatedData map[string]string
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedClient, err := service.UPDATE_CLIENT(id, updatedData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedClient})
}

func delete(c *gin.Context) {
	id := c.Param("id")
	_, err := service.DELETE_CLIENT(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted data"})
}

func getById(c *gin.Context) {
	id := c.Param("id")
	client, err := service.GET_CLIENT_ID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, client)
}
