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
			pet.POST("/", controller.CreatePet)
			pet.PUT("/:id", controller.UpdatePet)
			pet.DELETE("/:id", controller.DeletePet)
			pet.GET("/all", controller.GetAllPet)
			pet.GET("/:id", controller.GetByIdPet)
		}
	}

	router_bus.Run(":8080")
}
