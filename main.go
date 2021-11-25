package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"

	"app-share-api/infra"
	"app-share-api/interface/handler"
	"app-share-api/usecase"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 

	db := infra.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	postRepository := infra.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase)

	likeRepository := infra.NewLikeRepository(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepository)
	likeHandler := handler.NewLikeHandler(likeUsecase)

	commentRepository := infra.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	commentHandler := handler.NewCommentHandler(commentUsecase)

	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	handler.InitRouting(e, postHandler, likeHandler, commentHandler, userHandler, authHandler)
	e.Logger.Fatal(e.Start(":8080"))
}