package repository

import (
	"context"
	"errors"

	"github.com/cpkal/golang-user-service/api/models"
)

type UserRepository interface {
	GetByID(ctx context.Context, userID string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, userID string, user *models.User) error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User

	return &user, errors.New("New error")
}

func (r *UserRepositoryImpl) Create(ctx context.Context, User *models.User) error {
	return errors.New("New Error")
}

func (r *UserRepositoryImpl) Update(ctx context.Context, userID string, User *models.User) error {
	return errors.New("New error")
}
