package routes

import (
	"github.com/gin-gonic/gin"
	handler "obsidian/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.Login)
}
