package repositories

import "obsidian/models"

type UserDAO interface {
	CreateUser(user *models.User) (*models.User, error)
}
