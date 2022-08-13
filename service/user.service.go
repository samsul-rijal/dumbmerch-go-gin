package service

import (
	"dumbmerch-go/dto"
	"dumbmerch-go/entity"
	"dumbmerch-go/repository"
	"errors"
	"log"

	"github.com/mashingan/smapping"

	_user "dumbmerch-go/service/user"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) (*_user.UserResponse, error)
	UpdateUser(updateUserRequest dto.UpdateUserRequest) (*_user.UserResponse, error)
	FindUserByEmail(email string) (*_user.UserResponse, error)
	FindUserByID(userID string) (*_user.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) UpdateUser(updateUserRequest dto.UpdateUserRequest) (*_user.UserResponse, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&updateUserRequest))

	if err != nil {
		return nil, err
	}

	user, err = c.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	res := _user.NewUserResponse(user)
	return &res, nil

}

func (c *userService) CreateUser(registerRequest dto.RegisterRequest) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByEmail(registerRequest.Email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	user, _ = c.userRepo.InsertUser(user)
	res := _user.NewUserResponse(user)
	return &res, nil

}

func (c *userService) FindUserByEmail(email string) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	userResponse := _user.NewUserResponse(user)
	return &userResponse, nil
}

func (c *userService) FindUserByID(userID string) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	userResponse := _user.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}
