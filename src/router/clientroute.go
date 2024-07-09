package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func create() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/client", controller.Create)
		api.PUT("/client/:id", controller.Update)
		api.DELETE("/client/:id", controller.Delete)
		api.GET("/client/all", controller.GetAll)
		api.GET("/client/:id", controller.GetById)
	}

	router.Run(":8080")
}
