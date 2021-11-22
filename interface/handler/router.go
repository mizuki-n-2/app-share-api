package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, postHandler PostHandler) {
	e.GET("/posts", postHandler.GetPosts())
	e.GET("/posts/:id", postHandler.GetPost())
	e.POST("/posts", postHandler.CreatePost())
	e.PATCH("/posts/:id", postHandler.UpdatePost())
	e.DELETE("/posts/:id", postHandler.DeletePost())
}
