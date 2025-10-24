package services

import (
	"log"
	"obsidian/config"
	"obsidian/models"
	"obsidian/repositories"
	"obsidian/utils"
	"strings"
)

type UserService interface {
	CreateUSER(user *models.User) (*models.User, int, error)
	Login(email string, password string) (*config.TokenPair, int, error)
	FindUserById(id int) (*models.User, error)
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
	if user.Password == "" || user.Email == "" {
		return nil, 400, utils.BindError("email or password cannot be empty")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v", user.Email, err)
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

func (u *userService) Login(email string, password string) (*config.TokenPair, int, error) {
	user, err := u.repo.FindUserByEmail(email)

	if err != nil {
		return nil, 404, utils.BindError("User not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, 401, utils.BindError("Invalid password")
	}

	token, err := config.GenerateTokenPairs(user.ID)
	if err != nil {
		return nil, 500, nil
	}
	return token, 200, nil
}

func (u *userService) FindUserById(id int) (*models.User, error) {
	user, err := u.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
