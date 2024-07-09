package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func API_BUS() {
	router_bus := gin.Default()
	api := router_bus.Group("/api")
	{
		client := api.Group("/client")
		{
			client.POST("/", controller.Create)
			client.PUT("/:id", controller.Update)
			client.DELETE("/:id", controller.Delete)
			client.GET("/all", controller.GetAll)
			client.GET("/:id", controller.GetById)
		}
		pet := api.Group("/pet")
		{
			pet.POST("/pet", controller.CreatePet)
			pet.PUT("/pet/:id", controller.UpdatePet)
			pet.DELETE("/pet/:id", controller.DeletePet)
			pet.GET("/pet/all", controller.GetAllPet)
			pet.GET("/pet/:id", controller.GetByIdPet)
		}
	}

	router_bus.Run(":8080")
}
