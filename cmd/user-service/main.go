package main

import (
	"context"
	"net/http"
	
	"github.com/cpkal/golang-user-service/internal/repository"
	"github.com/cpkal/golang-user-service/internal/service"
	"github.com/cpkal/golang-user-service/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//db init
	DB.Run()

	//registering services
	UserRepository := repository.NewUserRepository()

	userService := service.NewUserService(UserRepository)

	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/auth/register", func(c *gin.Context) {	
		c.JSON(http.StatusOK, gin.H{
			"error": "asdsda",
		})
	})

	v1.GET("/auth/login", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "login",
		})
	})

	r.Run()
}
