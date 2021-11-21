package model

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImagePath string `json:"image_path"`
	AppURL    string `json:"app_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Posts []Post

func FindPosts() Posts {
	var posts Posts
	DB.Find(&posts)
	return posts
}

func CreatePost(post *Post) {
	DB.Create(post)
}
