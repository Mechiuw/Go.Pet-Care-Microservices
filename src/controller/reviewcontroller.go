package controller

import (
	"ginpet/src/model"
	"ginpet/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	review := model.Review{}

	if err := c.Bind(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createReview, err := service.CREATE_REVIEW(review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created review", "data": createReview,
	})
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	review := model.Review{}
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updateReview, err := service.UPDATE_REVIEW(id, review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated review",
		"data":    updateReview,
	})
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	deleteServiceReview, err := service.DELETE_REVIEW(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted review",
		"data":    deleteServiceReview,
	})
}

func GetReviewById(c *gin.Context) {
	id := c.Param("id")

	fetchReview, err := service.GET_BY_ID_REVIEW(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully fetch review",
		"data":    fetchReview,
	})
}

func GetAllReview(c *gin.Context) {
	fetchReview, err := service.GET_ALL_REVIEW()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully fetch review",
		"data":    fetchReview,
	})
}
