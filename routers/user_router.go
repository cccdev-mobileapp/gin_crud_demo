package routers

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	// router.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	// router.HandleFunc("/user", controllers.AddNewUser).Methods("POST")
	// router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")

	// router.GET("/user", controllers.GetAllUsers)

	userGroup := router.Group("/user")
	{
		userGroup.GET("/", controllers.GetAllUsers)
		userGroup.POST("/", controllers.AddNewUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
	}
}
