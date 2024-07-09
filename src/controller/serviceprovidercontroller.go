package controller

import (
	"ginpet/src/model"
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created service provider",
		"data":    sp,
	})
}
func UpdateSp(c *gin.Context)  {}
func DeleteSp(c *gin.Context)  {}
func GetAllSp(c *gin.Context)  {}
func GetByIdSp(c *gin.Context) {}
