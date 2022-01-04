package dto

import "time"

type Comment struct {
	ID         string    `json:"id"`
	PostID     string    `json:"post_id"`
	UserID     string    `json:"user_id"`
	Content    string    `json:"content"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserName   string    `json:"user_name"`
	UserAvatar string    `json:"user_avatar"`
	LikesCount int       `json:"likes_count"`
}
