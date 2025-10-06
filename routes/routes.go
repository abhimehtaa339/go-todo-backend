package routes

import (
	"github.com/gin-gonic/gin"
	"obsidian/controllers"
)

func ResisterRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.CreateUser)
}
