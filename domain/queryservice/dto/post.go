package dto

import (
	"time"
)

type Post struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Image         string    `json:"image"`
	AppURL        string    `json:"app_url"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserName      string    `json:"user_name"`
	UserAvatar    string    `json:"user_avatar"`
	LikesCount    int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
}
