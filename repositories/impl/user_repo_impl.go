package impl

import (
	"errors"
	"fmt"
	"obsidian/config"
	"obsidian/models"
	"obsidian/repositories/interfaces"
	"obsidian/utils"
	"strings"
)

type userDAO struct{}

func NewUserDAO() interfaces.dao {
	return &userDAO{}
}

func (u *userDAO) CreateUSER(user *models.User) (*models.User, error) {
	if passHash, err := utils.HashPassword(user.Password); err != nil {
		return nil, err
	} else {
		user.Password = passHash
	}

	if err := config.DB.Create(user).Error; err != nil {

		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errors.New(fmt.Sprint("exists"))
		}
		return nil, err
	}
	return user, nil
}
