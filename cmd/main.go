package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"obsidian/config"
	"obsidian/models"
	"obsidian/routes"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.RegisterRoutes(r)

	port := config.GetEnv("PORT", "8001")

	r.Run(fmt.Sprintf(":%s", port))
}
