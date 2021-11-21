package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app-share-api/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/posts", handler.GetPosts())
	e.GET("/posts/:id", handler.GetPost())
	e.POST("/posts", handler.CreatePost())
	e.PATCH("/posts/:id", handler.UpdatePost())
	e.DELETE("/posts/:id", handler.DeletePost())

	e.Logger.Fatal(e.Start(":1323"))
}