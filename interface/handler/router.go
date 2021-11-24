package handler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting(e *echo.Echo, postHandler PostHandler, likeHandler LikeHandler, userHandler UserHandler, authHandler AuthHandler) {
	api := e.Group("/api")
	// 認証なし
	api.POST("/users", userHandler.CreateUser())
	api.GET("/users/:id", userHandler.GetUser())
	api.POST("/login", authHandler.Login())
	api.GET("/posts", postHandler.GetAllPosts())
	api.GET("/posts/:id", postHandler.GetPost())

	// 認証あり
	auth := api.Group("/auth")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	auth.POST("/posts", postHandler.CreatePost())
	auth.PATCH("/posts/:id", postHandler.UpdatePost())
	auth.DELETE("/posts/:id", postHandler.DeletePost())
	auth.POST("/likes", likeHandler.Like())
	auth.DELETE("/likes/:id", likeHandler.Unlike())
}
