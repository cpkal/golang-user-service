package main

import (
	"context"
	"net/http"
	"time"

	"github.com/cpkal/golang-user-service/api/models"
	"github.com/cpkal/golang-user-service/internal/repository"
	"github.com/cpkal/golang-user-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//registering services
	UserRepository := repository.NewUserRepository()

	userService := service.NewUserService(UserRepository)

	userService.GetUserByID(context.Background(), "1")

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		var user models.User
		user.ID = "1"
		user.Email = "dummy@dummy.com"
		user.Password = "supersecret"
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		err := userService.CreateUser(c, &user)

		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
