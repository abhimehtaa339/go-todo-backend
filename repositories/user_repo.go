package repositories

import "obsidian/models"

type UserDAO interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id int) (*models.User, error)
}
