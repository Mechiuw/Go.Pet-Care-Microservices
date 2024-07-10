package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func API_BUS() {
	router_bus := gin.Default()
	api := router_bus.Group("/api")

	clientRoutes(api.Group("/client"))
	petRoutes(api.Group("/pet"))
	serviceProviderRoutes(api.Group("/service-provider"))
	appointmentRoutes(api.Group("/appointment"))
	reviewRoutes(api.Group("/review"))

	router_bus.Run(":8080")
}

func clientRoutes(client *gin.RouterGroup) {
	client.POST("/", controller.Create)
	client.PUT("/:id", controller.Update)
	client.DELETE("/:id", controller.Delete)
	client.GET("/all", controller.GetAll)
	client.GET("/:id", controller.GetById)
}

func petRoutes(pet *gin.RouterGroup) {
	pet.POST("/", controller.CreatePet)
	pet.PUT("/:id", controller.UpdatePet)
	pet.DELETE("/:id", controller.DeletePet)
	pet.GET("/all", controller.GetAllPet)
	pet.GET("/:id", controller.GetByIdPet)
}

func serviceProviderRoutes(sp *gin.RouterGroup) {
	sp.POST("/", controller.CreateSp)
	sp.PUT("/:id", controller.UpdateSp)
	sp.DELETE("/:id", controller.DeleteSp)
	sp.GET("/all", controller.GetAllSp)
	sp.GET("/:id", controller.GetByIdSp)
}

func appointmentRoutes(app *gin.RouterGroup) {
	app.POST("/", controller.CreateAppointment)
	app.PUT("/:id", controller.UpdateAppointment)
	app.DELETE("/:id", controller.DeleteAppointment)
	app.GET("/all", controller.GetAllAppointment)
	app.GET("/:id", controller.GetByIDAppointment)
}

func reviewRoutes(review *gin.RouterGroup) {
	review.POST("/", controller.CreateReview)
	review.PUT("/:id", controller.UpdateReview)
	review.DELETE("/:id", controller.DeleteReview)
	review.GET("/all", controller.GetAllReview)
	review.GET("/:id", controller.GetReviewById)
}
