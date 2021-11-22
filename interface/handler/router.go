package handler

import (
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo, postHandler PostHandler) {
	api := e.Group("/api")
	// api.GET("/posts", postHandler.GetPosts())
	// api.GET("/posts/:id", postHandler.GetPost())
	api.POST("/posts", postHandler.CreatePost())
	// api.PATCH("/posts/:id", postHandler.UpdatePost())
	// api.DELETE("/posts/:id", postHandler.DeletePost())
}
