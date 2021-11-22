package model

import (
	"errors"
)

type Post struct {
	ID        int   `json:"id"`
	UserID    int   `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// postのコンストラクタ
func NewPost(userID int, title, content string) (*Post, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	post := &Post{
		UserID:    userID,
		Title:     title,
		Content:   content,
	}

	return post, nil
}
