package controller

import (
	"ginpet/src/model"
	"ginpet/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAppointment(c *gin.Context) {
	var app model.Appointment

	if err := c.BindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdAppointment, err := service.CREATE_APPOINTMENT(app)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created appointment",
		"data":    createdAppointment,
	})
}

func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var app model.Appointment

	if err := c.BindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedAppointment, err := service.UPDATE_APPOINTMENT(id, app)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update appointment",
		"data":    updatedAppointment,
	})

}
