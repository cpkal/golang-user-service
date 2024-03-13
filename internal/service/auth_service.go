package service

import (
	"context"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	// jwt.StandardClaims
}


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
	var creds Credentials
	
	creds.Username = username
	creds.Password = password

	expirationTime := time.Now().Add(5 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		fmt.Println("Error in creating token")
		return "", err
	}

	return tokenString, nil
}

func (a *AuthServiceImpl) SignOutUser(ctx context.Context, token string) error {
	return nil
}
