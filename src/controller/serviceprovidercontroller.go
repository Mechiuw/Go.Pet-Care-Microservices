package controller

import (
	"ginpet/src/model"
	"ginpet/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSp(c *gin.Context) {
	var sp model.ServiceProvider

	if err := c.BindJSON(&sp); err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdSp, err := service.CREATE_SERVICE_PROVIDER(sp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created service provider",
		"data":    createdSp,
	})
}
func UpdateSp(c *gin.Context) {
	id := c.Param("id")
	var sp model.ServiceProvider
	if err := c.BindJSON(&sp); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedSp, err := service.UPDATE_SERVICE_PROVIDER(id, sp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated service provider",
		"data":    updatedSp,
	})

}

func DeleteSp(c *gin.Context) {
	id := c.Param("id")

	_, err := service.DELETE_SERVICE_PROVIDER(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted service provider"})
}
func GetAllSp(c *gin.Context) {
	fetch, err := service.GET_ALL_SERVICE_PROVIDER()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully fetch service provider", "data": fetch})
}
func GetByIdSp(c *gin.Context) {}