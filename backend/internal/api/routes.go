// internal/api/routes.go
package api

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the user endpoints.
func RegisterRoutes(router *gin.Engine, userController *UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:userId", userController.GetUserById)
		userRoutes.PUT("/:userId", userController.UpdateUser)
		userRoutes.DELETE("/:userId", userController.DeleteUser)
	}
}
