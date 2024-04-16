package api

import (
	"github.com/NavinduNavoda/waggle-gobackend/data"
	"github.com/NavinduNavoda/waggle-gobackend/server/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, db data.DBConfig) {

	router.POST("/api/user/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		// Perform your authentication logic here
		user, err := data.GetUserByUsername(username)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid username"})
			return
		}

		if user.Password != password {
			ctx.JSON(401, gin.H{"error": "Invalid password"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Logged in successfully",
			"sid": services.GetNewSessionId(),
			"db" : db,
		})
	})

	router.POST("/api/user/signup", func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		// Perform your registration logic here
		err := data.AddUser(data.User{Email: email, Username: username, Password: password})
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Failed to register user"})
			return
		}

		ctx.JSON(200, gin.H{"message": "User registered successfully"})
	})


	router.GET("/api/user/check", services.RequireAuth, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "User is authenticated"})
	})

	router.POST("/api/user/check", services.RequireAuth, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "User is authenticated"})
	})
}