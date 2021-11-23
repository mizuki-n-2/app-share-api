package handler

import (
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo, postHandler PostHandler, userHandler UserHandler, authHandler AuthHandler) {
	api := e.Group("/api")
	// 認証なし
	api.POST("/users", userHandler.CreateUser())
	api.GET("/users/:id", userHandler.GetUser())
	api.POST("/login", authHandler.Login())
	api.GET("/posts", postHandler.GetAllPosts())
	api.GET("/posts/:id", postHandler.GetPost())

	// 認証あり
	auth := api.Group("/auth")
	auth.POST("/posts", postHandler.CreatePost())
	auth.PATCH("/posts/:id", postHandler.UpdatePost())
	auth.DELETE("/posts/:id", postHandler.DeletePost())
}
