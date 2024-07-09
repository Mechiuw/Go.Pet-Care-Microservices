package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func PET() {
	// api path http://localhost:8080/api/pet

	router_pet := gin.Default()
	api := router_pet.Group("/api")
	{
		pet := api.Group("/pet")
		{
			pet.POST("/pet", controller.CreatePet)
			pet.PUT("/pet/:id", controller.UpdatePet)
			pet.DELETE("/pet/:id", controller.DeletePet)
			pet.GET("/pet/all", controller.GetAllPet)
			pet.GET("/pet/:id", controller.GetByIdPet)
		}
	}
	router_pet.Run(":8080")
}
