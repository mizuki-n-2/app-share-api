package infra

import (
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"app-share-api/domain/model"
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

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Post{}, model.User{}, model.Comment{}, model.Like{})

	return db
}

func GetDB() *gorm.DB {
	return db
}