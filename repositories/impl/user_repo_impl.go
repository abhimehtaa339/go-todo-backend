package impl

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"obsidian/config"
	"obsidian/models"
	"obsidian/repositories"
	"obsidian/utils"
)

type userDAO struct{}

func NewUserDAO() repositories.UserDAO {
	return &userDAO{}
}

func (u *userDAO) CreateUser(user *models.User) (*models.User, error) {
	err := config.DB.Create(user).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, utils.BindError("user already exists")
		}

		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, utils.BindError("invalid user data")
		}

		return nil, err
	}
	return user, nil
}

func (u *userDAO) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.BindError("invalid user data")
	}
	return &user, nil
}
