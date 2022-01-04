package router

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app-share-api/interface/handler"
)

func NewRouter(e *echo.Echo, postHandler handler.PostHandler, likeHandler handler.LikeHandler, commentHandler handler.CommentHandler, userHandler handler.UserHandler, authHandler handler.AuthHandler) {
	api := e.Group("/api")
	// 認証なし
	api.POST("/users", userHandler.CreateUser())
	api.GET("/users/:id", userHandler.GetUser())
	api.GET("/users/rank/like", userHandler.GetRankingLikeUsers())
	api.GET("/users/rank/post", userHandler.GetRankingPostUsers())
	api.POST("/login", authHandler.Login())
	api.GET("/posts", postHandler.GetPosts())
	api.GET("/posts/like", postHandler.GetLikePosts())
	api.GET("/posts/:id", postHandler.GetPost())
	api.GET("/likes", likeHandler.GetLikes())
	api.GET("/comments", commentHandler.GetComments())

	// 認証あり
	auth := api.Group("/auth")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	auth.PATCH("/users/:id", userHandler.UpdateUser())
	auth.POST("/users/:id/avatar", userHandler.UploadUserAvatar())
	auth.POST("/posts", postHandler.CreatePost())
	auth.PATCH("/posts/:id", postHandler.UpdatePost())
	auth.POST("/posts/:id/image", postHandler.UploadPostImage())
	auth.DELETE("/posts/:id", postHandler.DeletePost())
	auth.POST("/comments", commentHandler.CreateComment())
	auth.PATCH("/comments/:id", commentHandler.UpdateComment())
	auth.DELETE("/comments/:id", commentHandler.DeleteComment())
	auth.POST("/likes", likeHandler.Like())
	auth.DELETE("/likes/:id", likeHandler.Unlike())
}
