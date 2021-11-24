package handler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting(e *echo.Echo, postHandler PostHandler, likeHandler LikeHandler, commentHandler CommentHandler, userHandler UserHandler, authHandler AuthHandler) {
	api := e.Group("/api")
	// 認証なし
	api.POST("/users", userHandler.CreateUser())
	api.GET("/users/:id", userHandler.GetUser())
	api.POST("/login", authHandler.Login())
	api.GET("/posts", postHandler.GetAllPosts())
	api.GET("/posts/:id", postHandler.GetPost())
	api.GET("/posts/:post_id/comments", commentHandler.GetComments())

	// 認証あり
	auth := api.Group("/auth")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	auth.POST("/posts", postHandler.CreatePost())
	auth.PATCH("/posts/:id", postHandler.UpdatePost())
	auth.DELETE("/posts/:id", postHandler.DeletePost())
	auth.POST("/posts/:post_id/likes", likeHandler.Like())
	auth.DELETE("/posts/:post_id/likes/:id", likeHandler.Unlike())
	auth.POST("/posts/:post_id/comments", commentHandler.CreateComment())
	auth.PATCH("/posts/:post_id/comments/:id", commentHandler.UpdateComment())
	auth.DELETE("/posts/:post_id/comments/:id", commentHandler.DeleteComment())
}
