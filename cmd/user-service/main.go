package main

import (
	"context"

	"github.com/cpkal/golang-user-service/internal/repository"
	"github.com/cpkal/golang-user-service/internal/service"
)

func main() {
	//registering services
	UserRepository := repository.NewUserRepository()

	userService := service.NewUserService(UserRepository)

	userService.GetUserByID(context.Background(), "1")

	//startServer()
}
