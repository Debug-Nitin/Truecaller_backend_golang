package routes

import (
	"coding-task/controller"
	"coding-task/middleware"

	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func(r *gin.Engine) {
	r.POST("/register", controller.RegisterUser)

	r.Use(middleware.AuthMiddleware())
	r.POST("/mark-as-spam", controller.MarkAsSpam)
	r.POST("/search-by-phone", controller.SearchByPhone)
	r.POST("/search-by-name", controller.SearchByName)
}
