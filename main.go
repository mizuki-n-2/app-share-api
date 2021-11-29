package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app-share-api/infra"
	"app-share-api/infra/repositoryImpl"
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

	e.Static("/", "static")

	postRepository := repositoryImpl.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase)

	likeRepository := repositoryImpl.NewLikeRepository(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepository)
	likeHandler := handler.NewLikeHandler(likeUsecase)

	commentRepository := repositoryImpl.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	commentHandler := handler.NewCommentHandler(commentUsecase)

	userRepository := repositoryImpl.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	handler.InitRouting(e, postHandler, likeHandler, commentHandler, userHandler, authHandler)
	e.Logger.Fatal(e.Start(":8080"))
}