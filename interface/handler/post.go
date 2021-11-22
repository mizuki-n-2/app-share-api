package handler

import (
	"github.com/labstack/echo"
)

type PostHandler interface {
	GetPosts() echo.HandlerFunc
	GetPost() echo.HandlerFunc
	CreatePost() echo.HandlerFunc
	UpdatePost() echo.HandlerFunc
	DeletePost() echo.HandlerFunc
}
