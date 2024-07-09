package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func PET() {
	// api path http://localhost:8080/api/pet

	router := gin.Default()
	api := router.Group("/api")
	{
		pet := api.Group("/pet")
		{
			pet.POST("/", controller.CreatePet)
			pet.PUT("/", controller.UpdatePet)
			pet.DELETE("/:id", controller.DeletePet)
			pet.GET("/all", controller.GetAllPet)
			pet.GET("/:id", controller.GetByIdPet)
		}
	}

	router.Run(":8080")
}
