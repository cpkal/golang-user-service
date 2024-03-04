package service

import (
	"context"
)

type AuthService interface {
	SignInUser(ctx context.Context, username string, password string) (string, error)
	SignOutUser(ctx context.Context, token string) error
}

type AuthServiceImpl struct {
	// DB db.DB
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (a *AuthServiceImpl) SignInUser(ctx context.Context, username string, password string) (string, error) {
	
	return "", nil
}

func (a *AuthServiceImpl) SignOutUser(ctx context.Context, token string) error {
	return nil
}
