package handler

import (
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo, postHandler PostHandler) {
	// 認証あり
	api := e.Group("/api")
	// 投稿API
	api.GET("/posts", postHandler.GetAllPosts())
	api.GET("/posts/:id", postHandler.GetPost())
	api.POST("/posts", postHandler.CreatePost())
	api.PATCH("/posts/:id", postHandler.UpdatePost())
	api.DELETE("/posts/:id", postHandler.DeletePost())
}
