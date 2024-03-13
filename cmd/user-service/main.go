package main

import (
	"net/http"
	
	// "github.com/cpkal/golang-user-service/internal/repository"
	// "github.com/cpkal/golang-user-service/internal/service"
	"github.com/cpkal/golang-user-service/internal/middleware"
	"github.com/cpkal/golang-user-service/internal/service"
	"github.com/cpkal/golang-user-service/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//db init
	db, err := db.NewMySQLDB()

	if err != nil {
		panic(err)
	}

	defer db.Conn.Close()

	//registering services
	// UserRepository := repository.NewUserRepository()

	// userService := service.NewUserService(UserRepository)
	authService := service.NewAuthService()

	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.Use(middleware.tokenAuthMiddleware())

	v1.POST("/auth/register", func(c *gin.Context) {	
		// userService.CreateUser(c)
		c.JSON(http.StatusOK, gin.H{
			"error": "",
		})
	})

	v1.POST("/auth/login", func(c *gin.Context) {

		tokenString, err := authService.SignInUser(c, c.PostForm("username"), c.PostForm("password"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	v1.GET("/auth/user/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})

	r.Run()
}
