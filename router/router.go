package router

import (
	"task5-pbi-btpns-ArifBudiantoPratomo/controllers"
	"task5-pbi-btpns-ArifBudiantoPratomo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.POST("/api/users/register", controllers.Register)
	router.POST("/api/users/login", controllers.Login)

	router.Use(middleware.Authentication)

	users := router.Group("/api/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUserById)
		users.PUT("/:id", middleware.AuthUser, controllers.UpdateUser)
		users.DELETE("/:id", middleware.AuthUser, controllers.DeleteUser)
	}

	photos := router.Group("/api/photos")
	{
		photos.GET("/", controllers.GetPhotos)
		photos.GET("/:id", controllers.GetPhotoById)
		photos.POST("/", controllers.CreatePhoto)
		photos.PUT("/:id", middleware.AuthPhoto, controllers.UpdatePhoto)
		photos.DELETE("/:id", middleware.AuthPhoto, controllers.DeletePhoto)
	}

	return router
}
