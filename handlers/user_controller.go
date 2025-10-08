package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"obsidian/models"
	"obsidian/repositories/impl"
	"obsidian/services"
)

var userService = services.NewUserService(impl.NewUserDAO())

func SignUp(c *gin.Context) {
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
		"username":   newUser.Username,
		"email":      newUser.Email,
		"name":       fmt.Sprintf("%s %s", newUser.FirstName, newUser.LastName),
		"dateJoined": newUser.DateJoined,
	})
}
