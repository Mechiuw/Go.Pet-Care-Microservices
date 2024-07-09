package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func CLIENT() {
	router_client := gin.Default()

	api := router_client.Group("/api")
	{
		client := api.Group("/client")
		{
			client.POST("/", controller.Create)
			client.PUT("/:id", controller.Update)
			client.DELETE("/:id", controller.Delete)
			client.GET("/all", controller.GetAll)
			client.GET("/:id", controller.GetById)
		}
	}

	router_client.Run(":8080")
}
