package infra

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	pm "app-share-api/domain/model/post"
	um "app-share-api/domain/model/user"
	cm "app-share-api/domain/model/comment"
	lm "app-share-api/domain/model/like"
)

var db *gorm.DB
var err error

func InitDB() *gorm.DB {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(pm.Post{}, um.User{}, cm.Comment{}, lm.Like{})

	return db
}

func GetDB() *gorm.DB {
	return db
}