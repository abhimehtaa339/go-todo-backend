package services

import (
	"log"
	"obsidian/models"
	"obsidian/repositories"
	"obsidian/utils"
	"strings"
)

type UserService interface {
	CreateUSER(user *models.User) (*models.User, int, error)
}

type userService struct {
	repo repositories.UserDAO
}

func NewUserService(repo repositories.UserDAO) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) CreateUSER(user *models.User) (*models.User, int, error) {

	//validations
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return nil, 400, utils.BindError("username, password, or email cannot be empty")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v", user.Username, err)
		return nil, 500, utils.BindError("Something went wrong, please try again")
	}
	user.Password = hashed

	newUser, err := u.repo.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			return nil, 409, err
		}
		if strings.Contains(err.Error(), "invalid user data") {
			return nil, 400, utils.BindError("invalid user data")
		}
		return nil, 500, utils.BindError("Something went wrong")
	}
	return newUser, 201, nil
}
