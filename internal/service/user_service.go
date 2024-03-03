package service

import (
	"context"
	"errors"

	"github.com/cpkal/golang-user-service/api/models"
	"github.com/cpkal/golang-user-service/internal/repository"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, userID string, user *models.User) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: repo,
	}
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User

	return &user, errors.New("error")
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, User *models.User) error {
	return errors.New("new error")
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, userID string, User *models.User) error {
	return errors.New("err")
}
