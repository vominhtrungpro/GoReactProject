package rest

import (
	"GoReactProject/internal/handler/rest/authenticationapi"
	"GoReactProject/internal/handler/rest/avatarapi"
	"GoReactProject/internal/handler/rest/crawldataapi"
	"GoReactProject/internal/handler/rest/midlewareapi"
	"GoReactProject/internal/handler/rest/userapi"

	"github.com/gin-gonic/gin"
)

func ProjectAPI() {
	router := gin.Default()
	router.Use(midlewareapi.CORSMiddleware())
	router.GET("/user/all", userapi.GetUsers)
	router.GET("/user/:id", userapi.GetUserById)
	router.GET("/user/email/:email", userapi.GetUserByEmail)
	router.POST("/user", userapi.CreateUser)
	router.PUT("/user", userapi.UpdateUser)
	router.DELETE("/user", userapi.DeleteUser)
	router.POST("/login", authenticationapi.Login)
	router.GET("/login-google/:token", authenticationapi.LoginGoogle)
	router.GET("/auth", midlewareapi.RequireAuth)
	router.POST("/avatar/:id", avatarapi.CreateAvatar)
	router.GET("/avatar/:id", avatarapi.GetAvatar)
	router.POST("/crawl", crawldataapi.CrawlData)
	router.Run("localhost:9090")
}
