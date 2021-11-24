package model

import (
	"errors"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPost(userID int, title, content string) (*Post, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	post := &Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	return post, nil
}

func (post *Post) Update(title, content string) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	post.Title = title
	post.Content = content

	return nil
}
