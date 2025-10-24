package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"obsidian/models"
	"obsidian/services"
)

func SignUp(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser, status, err := userService.CreateUSER(&user)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}

		c.JSON(status, gin.H{
			"email":      newUser.Email,
			"name":       fmt.Sprintf("%s %s", newUser.FirstName, newUser.LastName),
			"dateJoined": newUser.DateJoined,
		})
	}
}

func Login(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, status, err := userService.Login(request.Email, request.Password)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}

		c.JSON(status, token)
	}
}

func ProtectedRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"cool": true})
	}
}
