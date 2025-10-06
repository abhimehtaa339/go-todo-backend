package dao

import "obsidian/models"

type UserDAO interface {
	CreateUSER(user *models.User) (*models.User, error)
}
