package routes

import (
	"github.com/gin-gonic/gin"
	handler "obsidian/handlers"
	"obsidian/middlewares"
	"obsidian/services"
)

func RegisterRoutes(r *gin.Engine, userService services.UserService) {
	r.POST("/signup", handler.SignUp(userService))
	r.POST("/login", handler.Login(userService))

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleWare(userService))
	{
		protected.GET("/profile", handler.ProtectedRoute())
	}
}
