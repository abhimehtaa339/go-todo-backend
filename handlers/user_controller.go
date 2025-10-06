package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"obsidian/models"
	"obsidian/repositories/impl"
)

var userDao = impl.NewUserDAO()

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := userDao.CreateUSER(&user)
	if err != nil {
		if err.Error() == "exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "User with this email or username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
