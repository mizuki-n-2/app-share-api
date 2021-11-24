package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	"app-share-api/domain/model/user"
	"app-share-api/domain/model/post"
	"app-share-api/domain/model/like"
	"app-share-api/domain/model/comment"
	"app-share-api/infra"
	"app-share-api/interface/handler"
	"app-share-api/usecase"
)

var db *gorm.DB
var err error

func initDB() *gorm.DB {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	)

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(post.Post{})
	db.AutoMigrate(user.User{})
	db.AutoMigrate(like.Like{})
	db.AutoMigrate(comment.Comment{})

	return db
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 

	db = initDB()

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