package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"obsidian/config"
	"obsidian/models"
	"obsidian/repositories"
	"obsidian/repositories/impl"
	"obsidian/routes"
	"obsidian/services"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	// Initialize DAO (concrete implementation of UserDAO)
	var userRepo repositories.UserDAO = impl.NewUserDAO(config.DB)

	userService := services.NewUserService(userRepo)

	r := gin.Default()

	routes.RegisterRoutes(r, userService)

	port := config.GetEnv("PORT", "8001")

	r.Run(fmt.Sprintf(":%s", port))
}
